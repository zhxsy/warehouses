package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go.uber.org/atomic"
	"log"
	"os"
	"time"
)

var connection = &TMQConnection{}

type TMQMessage struct {
	Key string
	Msg interface{}
}

// TMQConnection mq connection
type TMQConnection struct {
	Conn         *amqp.Connection
	Consumers    map[string]*TConsumer
	Logger       Logger
	DelayWrapper DelayWrapper
	CloseChan    chan bool
}

type Logger interface {
	Print(Val ...interface{})
	Fatal(val ...interface{})
}

func InitMQ(url string, vhost string, l Logger, delayWrapper DelayWrapper, publishFlag bool) *TMQConnection {
	connection.Logger = log.New(os.Stdout, "", log.LstdFlags)
	connection.Consumers = make(map[string]*TConsumer)
	connection.DelayWrapper = delayWrapper
	connection.CloseChan = make(chan bool)

	if l != nil {
		connection.Logger = l
	}
	var err error
	connection.Conn, err = amqp.DialConfig(url, amqp.Config{Vhost: vhost})
	//connection.Conn, err = amqp.Dial(url)
	if err != nil {
		connection.Logger.Fatal(err)
	}
	if publishFlag {
		go connection.RunDelayConsumer()
	}
	return connection
}

func GetConnection() *TMQConnection {
	return connection
}

func (c *TMQConnection) ConsumerDeclare(exchange, key string, tc *TConsumer) {
	tc.Conn = c.Conn
	ch, _ := c.GetChannel()

	c.Consumers[tc.Consumer] = tc
	//tc.Channel = ch
	ch.QueueDeclare(tc.Queue, true, false, tc.Exclusive, tc.NoWait, tc.Args)
	ch.QueueBind(tc.Queue, key, exchange, tc.NoWait, tc.Args)

}

func (c *TMQConnection) ConsumerRegister(exchange, key string, tc *TConsumer, fn func(delivery amqp.Delivery) error) {
	tc.Conn = c.Conn
	ch, err := c.GetChannel()
	if err != nil {
		c.Logger.Print(err.Error())
	}
	defer ch.Close()
	c.Consumers[tc.Consumer] = tc

	ch.QueueDeclare(tc.Queue, true, false, tc.Exclusive, tc.NoWait, tc.Args)
	ch.QueueBind(tc.Queue, key, exchange, tc.NoWait, tc.Args)

	go tc.Loop(fn)
}

//func (c *TMQConnection) Qos(tc *TConsumer, prefetchCount, prefetchSize int, flag bool) {
//	tc, ok := c.Consumers[tc.Consumer]
//	if !ok {
//		c.Logger.Print("please register consume")
//		return
//	}
//	err := tc.Channel.Qos(prefetchCount, prefetchSize, flag)
//	if err != nil {
//		c.Logger.Print(err)
//		defer tc.Channel.Close()
//	}
//
//}

func (c *TMQConnection) PublisherRegister(name string, t string) {
	if t == "" {
		t = amqp.ExchangeFanout
	}
	p := &TPublisher{
		IsLoop:    atomic.NewBool(false),
		Exchange:  name,
		Mandatory: false,
		Immediate: false,
		MsgCh:     make(chan TMQMessage),
	}
	ch, _ := c.GetChannel()
	var args = amqp.Table{}
	ch.ExchangeDeclare(
		name,
		t,
		true,
		false,
		false,
		false,
		args,
	)
	p.Conn = c.Conn
}

func (c *TMQConnection) DelayPublisherRegister(name string, t string) {
	if t == "" {
		t = amqp.ExchangeFanout
	}
	p := &TPublisher{
		IsLoop:    atomic.NewBool(false),
		Exchange:  name,
		Mandatory: false,
		Immediate: false,
		MsgCh:     make(chan TMQMessage),
	}
	ch, _ := c.GetChannel()
	var args = amqp.Table{}
	ch.ExchangeDeclare(
		name,
		t,
		true,
		false,
		false,
		false,
		args,
	)
	p.Conn = c.Conn
}

func (c *TMQConnection) Publish(name string, key string, msg interface{}) {
	b, err := json.Marshal(msg)
	if err != nil {
		c.Logger.Print("json marshal err: " + err.Error() + "\n")
		return
	}
	c.Logger.Print("ready to publish: " + string(b) + "\n")
	ch, err := c.Conn.Channel()
	if err != nil {
		c.DelayWrapper.SetDelayMsg(&DelayMsg{
			Name: name,
			Key:  key,
			Msg:  msg,
		})
		return
	}
	defer ch.Close()
	err = ch.Publish(name, key, false, false, amqp.Publishing{Body: b})
	if err != nil {
		c.Logger.Print("failed publish into exchange "+name+" key "+key+" msg ", msg, " err = "+err.Error()+"\n")
	}
}

func (c *TMQConnection) PublishDelay(name string, key string, msg interface{}, delay int) {
	err := c.DelayWrapper.SetDelayMsg(&DelayMsg{
		Name:      name,
		Key:       key,
		Msg:       msg,
		Delay:     int64(delay),
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		c.Logger.Print("publish delay failed: "+err.Error(), " message: ", msg)
	}
}

func (c *TMQConnection) RunDelayConsumer() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover err from delay consumer: %v", err)
			go c.RunDelayConsumer()
		}
	}()

	var ticker = time.NewTicker(time.Second)
	for {
		select {
		case <-c.CloseChan:
			c.Logger.Print("exit run delay consumer with close chan")
			break
		case <-ticker.C:
			msg, err := c.DelayWrapper.GetDelayMsg()
			if err != nil {
				c.Logger.Print("get delay msg failed: %v", err)
				continue
			}
			if msg == nil {
				continue
			}
			for i := range msg {
				c.Publish(msg[i].GetExchangeName(), msg[i].GetKey(), msg[i].GetMsg())
			}
		}
	}
}

func (c *TMQConnection) CLose() {
	c.CloseChan <- true
}

func (c *TMQConnection) GetChannel() (*amqp.Channel, error) {
	return c.Conn.Channel()
}
