package main

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

var rabbitmqUser = os.Getenv("RABBITMQ_USER")
var rabbitmqPassword = os.Getenv("RABBITMQ_PASSWORD")
var rabbitmqHost = os.Getenv("RABBITMQ_HOST")
var rabbitmqPort = os.Getenv("RABBITMQ_PORT")

func main() {
	// url := fmt.Sprintf("amqp://%s:%s@%s:%s", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort)
	url := "amqp://test:test@0.0.0.0:5672"

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("%s: %s(url: %s)", "failed to connect to rabbitmq", err, url)
	}
	defer conn.Close()

	chanel, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}
	defer chanel.Close()

	err = chanel.ExchangeDeclare(
		"topic-exchange", // name
		"topic",          // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a exchange", err)
	}

	queue, err := chanel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	for _, s := range os.Args[1:] {
		err = chanel.QueueBind(
			queue.Name,       // queue name
			s,                // routing key
			"topic-exchange", // exchange
			false,
			nil,
		)
	}

	if err != nil {
		log.Fatalf("%s: %s", "Failed to bind a queue", err)
	}
	msgs, err := chanel.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
