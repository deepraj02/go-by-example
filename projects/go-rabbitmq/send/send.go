package main

import (
	"bufio"
	"fmt"
	helper "go-by-example/projects/go-rabbitmq/helpers"
	"log"
	"os"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
	"golang.org/x/net/context"
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
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send: ")
		body, _ := reader.ReadString('\n')

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		err = ch.PublishWithContext(ctx, "", q.Name, false, false, ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		cancel()
		if err != nil {
			log.Printf("Failed to publish a message: %s", err)
		} else {
			log.Printf(" [x] Sent %s", body)
		}
	}
}
