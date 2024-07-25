package main

import (
	helper "go-by-example/projects/go-rabbitmq/helpers"
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	helper.FailOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	helper.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	helper.FailOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
