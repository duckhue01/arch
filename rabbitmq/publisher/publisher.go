package main

import (
	"log"
	"os"
	"strings"

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
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	err = chanel.Publish(
		"topic-exchange", // exchange
		os.Args[1],       // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyFrom(os.Args)),
		})
	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish message", err)
	}
	log.Print("message is sent!")

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
