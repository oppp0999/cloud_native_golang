package main

import (
	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic("not establish" + err.Error())
	}
	defer connection.Close()
}
