package main

import (
	"github.com/cfx/warehouses/library/mq"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	c := mq.InitMQ(mq.AmqpURL, "", nil, nil, false)

	c.ConsumerRegister(mq.DelayExchange, mq.Key, mq.DefaultConsumer(mq.DelayQueue, mq.DelayQueue, mq.Qos), Consume)
	c.ConsumerRegister(mq.DelayExchange, mq.Key, mq.DefaultConsumer(mq.DelayQueue1, mq.DelayQueue1, mq.Qos), Consume)
	c.ConsumerRegister(mq.NormalExchange, mq.Key, mq.DefaultConsumer(mq.Queue, mq.Queue, mq.Qos), Consume)

	var ch = make(chan int)
	ch <- 1
}

func Consume(d amqp.Delivery) error {
	log.Printf("%v:%d:%s", d.Exchange, time.Now().Unix(), string(d.Body))
	return nil
}
