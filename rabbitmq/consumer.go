package rabbitmq

import (
	"github.com/streadway/amqp"
)

func (r *RabbitMQ) DeclareQueue(queueName string) error {
	_, err := r.channel.QueueDeclare(
		queueName, // tên queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Bind queue với exchange
	err = r.channel.QueueBind(
		queueName,               // queue name
		"#",                     // routing key (nhận tất cả message)
		"notification_exchange", // exchange
		false,
		nil,
	)
	return err
}

func (r *RabbitMQ) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}
