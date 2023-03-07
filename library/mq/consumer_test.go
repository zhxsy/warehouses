package mq

import (
	"github.com/streadway/amqp"
	"log"
	"testing"
)

func init() {
	amqpUrl := "amqp://guest:guest@127.0.0.1:5672/"
	InitMQ(amqpUrl, "/dev", nil, nil, false)
}

func TestNewConsumer(t *testing.T) {
	consumner := NewConsumer("test_delayed_queue", "test_delayed_queue", false, false, false, false, nil)
	GetConnection().ConsumerRegister("test_delayed", "#", consumner, func(d amqp.Delivery) error {
		log.Print(string(d.Body))
		return nil
	})
	select {}
}
