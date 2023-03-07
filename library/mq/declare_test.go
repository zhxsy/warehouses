package mq

import (
	"github.com/streadway/amqp"
	"testing"
	"time"
)

func TestDeclareExchange(t *testing.T) {
	DeclareExchange("test_lizt", amqp.ExchangeFanout, false, false, false, false, nil)
}

func TestDeclareQueue(t *testing.T) {
	DeclareQueue("test_delayed", "test_delayed_queue", "*", false, false, false, false, nil)
}

func TestDeclareDelayExchange(t *testing.T) {
	DeclareExchange("test_delayed", "x-delayed-message", false, false, false, false, amqp.Table{
		"x-delayed-type": "topic",
	})
}

func TestPublishDelay(t *testing.T) {
	ch, _ := GetConnection().Conn.Channel()
	ch.Publish("test_delayed", "tag", false, false, amqp.Publishing{
		Headers:     map[string]interface{}{"x-delay": "10000"},
		ContentType: "test/plain",
		Body:        []byte("this is test 10: " + time.Now().String()),
	})
}
