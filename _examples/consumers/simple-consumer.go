package consumers

import (
	"fmt"
	"github.com/rozturac/rmqc"
	"github.com/rozturac/rmqc/_examples/events"
)

type SimpleConsumer struct {
}

func NewSimpleConsumer() *SimpleConsumer {
	return &SimpleConsumer{}
}

func (s SimpleConsumer) Configure(builder *rmqc.ConsumerBuilder) {
	builder.BindQueue("rabbitmq-test-queue")
	builder.SubscribeAsDirect("rabbitmq-test-exchange", "1")
	builder.SetConsumerName("SimpleConsumer")
	builder.SetPrefetchCount(4)
	builder.SetConsumerCount(4)
}

func (s SimpleConsumer) Consume(context *rmqc.ConsumerContext) {
	var event events.SimpleEvent
	context.Unmarshal(&event)
	fmt.Println(event)
}
