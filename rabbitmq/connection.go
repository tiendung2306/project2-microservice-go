package rabbitmq

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	// Get RabbitMQ connection details from environment variables
	host := os.Getenv("RABBITMQ_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("RABBITMQ_PORT")
	if port == "" {
		port = "5672"
	}
	user := os.Getenv("RABBITMQ_USER")
	if user == "" {
		user = "guest"
	}
	password := os.Getenv("RABBITMQ_PASSWORD")
	if password == "" {
		password = "guest"
	}

	// Construct connection URL
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	// Connect to RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %v", err)
	}

	// Declare exchange
	err = ch.ExchangeDeclare(
		"notification_exchange", // exchange name
		"topic",                 // exchange type
		true,                    // durable
		false,                   // auto-deleted
		false,                   // internal
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %v", err)
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
