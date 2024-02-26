package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var log = logrus.New()

func failOnError(err error, msg string) {
	//code..
}

func publishMessage(ch *amqp.Channel, q amqp.Queue, message string) {
	//code..
}

func handler(ch *amqp.Channel, q amqp.Queue) http.HandlerFunc {
	//code..
}

func main() {
	//code..
}
