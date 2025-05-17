package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type NotificationMessage struct {
	Type    string `json:"type"`   // "task", "auth", "user"
	Action  string `json:"action"` // "create", "update", "delete"
	UserID  uint   `json:"user_id"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (r *RabbitMQ) PublishMessage(message NotificationMessage) error {
	// Chuyển đổi message thành JSON
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Xác định routing key dựa trên loại message
	routingKey := message.Type + "." + message.Action

	// Publish message
	err = r.channel.Publish(
		"notification_exchange", // exchange
		routingKey,              // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	fmt.Println("Published message to RabbitMQ:", string(body))

	return err
}
