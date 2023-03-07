package mq

import (
	"github.com/streadway/amqp"
	"go.uber.org/atomic"
	"log"
	"time"
)

// TConsumer 消费者模型
type TConsumer struct {
	Conn      *amqp.Connection
	IsLoop    *atomic.Bool
	Exchange  string
	Key       string
	Consumer  string
	Queue     string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	Exit      chan bool
	Qos       int
}

func DefaultConsumer(name string, queue string, qos int) *TConsumer {
	return &TConsumer{
		Consumer:  name,
		Queue:     queue,
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
		Exit:      make(chan bool),
		Qos:       qos,
	}
}

func (c *TConsumer) SetAutoAck(arg bool) {
	c.AutoAck = arg
}

func (c *TConsumer) SetExclusive(arg bool) {
	c.Exclusive = arg
}

func (c *TConsumer) SetNoLocal(arg bool) {
	c.NoLocal = arg
}

func (c *TConsumer) SetNoWait(arg bool) {
	c.NoWait = arg
}

func (c *TConsumer) SetArgs(arg amqp.Table) {
	c.Args = arg
}

func (c *TConsumer) Loop(f func(amqp.Delivery) error) {
	// 暂时只支持手动 ACK
	if c.AutoAck == true {
		// TODO: auto ack
		return
	}

	ch, err := c.Conn.Channel()
	if err != nil {
		return
	}
	defer func() {
		if rErr := recover(); err != nil {
			log.Printf("ERROR: consumer %v exit with err: %v", c.Consumer, rErr)
		}
		ch.Close()
		log.Printf("ready to restart consumer %v", c.Consumer)
		go c.Loop(f)
	}()
	if c.Qos != 0 {
		ch.Qos(c.Qos, 0, false)
	}
	deliver, err := ch.Consume(c.Queue, c.Consumer, c.AutoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args)
	if err != nil {
		GetConnection().Logger.Print("get deliver failed, err: ", err)
		return
	}
	for {
		select {
		case msg := <-deliver:
			if err := f(msg); err != nil {
				log.Printf("ERROR: consumer %v deal msg failed: %+v", c.Consumer, err)
				time.Sleep(time.Duration(1) * time.Second) // 1秒在重试
				msg.Reject(true)
				continue
			}
			msg.Ack(false)
		case <-c.Exit:
			return
		}
	}
}

// NewConsumer .
func NewConsumer(name string, queue string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) *TConsumer {
	return &TConsumer{
		Consumer:  name,
		Queue:     queue,
		AutoAck:   autoAck,
		Exclusive: exclusive,
		NoLocal:   noLocal,
		NoWait:    noWait,
		Args:      args,
		Exit:      make(chan bool),
	}
}

func (c *TConsumer) StopServe() {
	c.Exit <- true
}
