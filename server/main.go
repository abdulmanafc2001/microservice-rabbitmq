package main

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Welcome to rabbitmq project")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	log.Println("successfully connected to rabbitmq server!")

	cha, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	que, err := cha.QueueDeclare(
		"test", // name
		false,  //durable
		false,  //auto delete
		false,  //exclusive
		false,  // nowait
		nil,    // arguments
	)

	if err != nil {
		log.Fatal(err)
	}

	err = cha.PublishWithContext(
		context.Background(),
		"",
		"test",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Test message"),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Queue status:", que)
	log.Println("Successfully published message")
}
