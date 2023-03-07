package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go.uber.org/atomic"
	"log"
	"strconv"
	"time"
)

type TPublisher struct {
	Conn      *amqp.Connection
	IsLoop    *atomic.Bool
	Exchange  string
	Mandatory bool
	Immediate bool
	MsgCh     chan TMQMessage
}

func NewPublisher(exchange string) *TPublisher {
	return &TPublisher{
		Exchange:  exchange,
		Mandatory: false,
		Immediate: false,
	}
}

func (p *TPublisher) Loop() {
	if p.IsLoop.Swap(true) {
		return
	}
	ch, _ := p.Conn.Channel()
	for {
		select {
		case msg := <-p.MsgCh:
			b, err := json.Marshal(msg.Msg)
			if err != nil {
				log.Printf("json marshal err: %v", err)
				return
			}
			log.Printf("ready to publish: %v", string(b))
			err = ch.Publish(p.Exchange, msg.Key, p.Mandatory, p.Immediate, amqp.Publishing{Body: b})
			if err != nil {
				log.Printf("publish err: %v", err)
				p.MsgCh <- msg
				return
			}
		case <-time.After(time.Minute):
			p.IsLoop.Swap(false)
			ch.Close()
			return
		}
	}
}

func (p *TPublisher) DelayLoop(delay int) {
	if p.IsLoop.Swap(true) {
		return
	}
	ch, _ := p.Conn.Channel()
	for {
		select {
		case msg := <-p.MsgCh:
			b, _ := json.Marshal(&msg.Msg)
			log.Printf("ready to delay publish: %v", string(b))
			err := ch.Publish("", p.Exchange+"_delay", p.Mandatory, p.Immediate, amqp.Publishing{Body: b, Expiration: strconv.Itoa(delay)})
			if err != nil {
				log.Printf("publish failed: %v", err)
			}
		case <-time.After(time.Minute):
			p.IsLoop.Swap(false)
			return
		}
	}
}
