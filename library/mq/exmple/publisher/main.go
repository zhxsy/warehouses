package main

import (
	"github.com/cfx/warehouses/library/mq"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	c := mq.InitMQ(mq.AmqpURL, "", nil, nil, false)

	c.DelayPublisherRegister(mq.DelayExchange, amqp.ExchangeFanout)

	go func(c *mq.TMQConnection) {
		for {
			c.PublishDelay(mq.DelayExchange, mq.Key, time.Now().Unix(), 10000)
			<-time.After(time.Second * 10)
		}
	}(c)

	var ch = make(chan int)
	ch <- 1
}
