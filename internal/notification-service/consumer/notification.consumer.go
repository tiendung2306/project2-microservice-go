package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"project2-microservice-go/internal/notification-service/dto"
	"project2-microservice-go/internal/notification-service/service"
	"project2-microservice-go/rabbitmq"
)

type NotificationConsumer struct {
	notificationService service.INotificationService
	rabbitmq            *rabbitmq.RabbitMQ
}

func NewNotificationConsumer(notificationService service.INotificationService, rabbitmq *rabbitmq.RabbitMQ) *NotificationConsumer {
	return &NotificationConsumer{
		notificationService: notificationService,
		rabbitmq:            rabbitmq,
	}
}

func (nc *NotificationConsumer) Start() error {
	// Declare queue
	queueName := "notification_queue"
	err := nc.rabbitmq.DeclareQueue(queueName)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}
	log.Printf("Successfully declared queue: %s", queueName)

	// Start consuming messages
	msgs, err := nc.rabbitmq.ConsumeMessages(queueName)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %v", err)
	}
	log.Println("Start consuming messages")

	// Process messages in a goroutine
	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", string(msg.Body))

			// Parse message
			var notificationMsg rabbitmq.NotificationMessage
			if err := json.Unmarshal(msg.Body, &notificationMsg); err != nil {
				log.Printf("Error parsing message: %v", err)
				msg.Ack(false)
				continue
			}
			log.Printf("Parsed message: %+v", notificationMsg)

			// Create notification record
			notificationReq := &dto.CreateNotificationRequest{
				UserID:     notificationMsg.UserID,
				Email:      notificationMsg.Email,
				Subject:    notificationMsg.Subject,
				Content:    notificationMsg.Content,
				NotifyType: notificationMsg.Type,
			}

			notification, err := nc.notificationService.CreateNotification(notificationReq)
			if err != nil {
				log.Printf("Error creating notification: %v", err)
				msg.Ack(false)
				continue
			}
			log.Printf("Created notification record: %+v", notification)

			// Send email
			emailReq := &dto.SendEmailRequest{
				To:      notificationMsg.Email,
				Subject: notificationMsg.Subject,
				Body:    notificationMsg.Content,
			}

			err = nc.notificationService.SendEmail(emailReq)
			if err != nil {
				log.Printf("Error sending email: %v", err)
				// Update notification status to failed
				_ = nc.notificationService.UpdateNotificationStatus(notification.ID, "failed")
				msg.Ack(false)
				continue
			}
			log.Printf("Successfully sent email to: %s", notificationMsg.Email)

			// Update notification status to sent
			err = nc.notificationService.UpdateNotificationStatus(notification.ID, "sent")
			if err != nil {
				log.Printf("Error updating notification status: %v", err)
			} else {
				log.Printf("Updated notification status to 'sent' for ID: %d", notification.ID)
			}

			// Acknowledge message
			msg.Ack(false)
			log.Printf("Message acknowledged")
		}
	}()

	return nil
}
