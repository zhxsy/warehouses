package mq

import (
	"github.com/streadway/amqp"
	"log"
)

func DeclareExchange(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	ch, err := GetConnection().GetChannel()
	if err != nil {
		log.Fatalf("get channel when declare exchange failed, err: %v", err)
	}

	err = ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
	if err != nil {
		log.Fatalf("declare exchange failed, err: %v", err)
	}
}

// DeclareQueue exchangeName 指向需要绑定的 exchange， queueName 需要创建的 queue
func DeclareQueue(exchangeName, queueName, key string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) {
	ch, err := GetConnection().GetChannel()
	if err != nil {
		log.Fatalf("get channel when declare queue failed, err: %v", err)
	}

	_, err = ch.QueueDeclare(queueName, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		log.Fatalf("declare queue failed, err: %v", err)
	}

	err = ch.QueueBind(queueName, key, exchangeName, noWait, args)
	if err != nil {
		log.Fatalf("bind queue to exhange failed, err: %v", err)
	}
}
