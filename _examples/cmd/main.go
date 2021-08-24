package main

import (
	"github.com/rozturac/rmqc"
	consumers "github.com/rozturac/rmqc/_examples/consumers"
	"github.com/rozturac/rmqc/_examples/events"
)

func main() {
	rbt, err := rmqc.Connect(rmqc.RabbitMQConfig{
		Host:           "localhost",
		Username:       "guest",
		Password:       "guest",
		Port:           "5672",
		VHost:          "/",
		ConnectionName: "rabbitmq-test",
		Reconnect: rmqc.Reconnect{
			MaxAttempt: 5,
			Interval:   100,
		},
	})
	if err != nil {
		panic(err)
	}

	consumer := consumers.NewSimpleConsumer()
	if err = rbt.BindConsumer(consumer); err != nil {
		panic(err)
	}
	rbt.Start()

	event := events.NewSimpleEvent("Hello World!")
	rbt.Publish("rabbitmq-test-exchange", "1", event)

	forever := make(chan bool)
	<-forever
}
