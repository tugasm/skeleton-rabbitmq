package main

import (
	"log"

	"github.com/streadway/amqp"
	gomail "gopkg.in/gomail.v2"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendMail(email, subject, content string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "example@hacktiv8.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(
		"test",
		0000,
		"test",
		"test",
	)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email sent to:", email)
	}
}

func SendSuccessCreateRent(email string) {
	SendMail(
		email,
		"Email Creation Successful",
		"<h1>Rent Created Successfully</h1><p>Your rent has been successfully created.</p>",
	)
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"emailQueue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			email := string(d.Body)
			log.Printf("Received a message: %s", email)
			// Mengirim email
			SendSuccessCreateRent(email)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
