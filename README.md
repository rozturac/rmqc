# rmqc
/Golang - RabbitMQ Connector v1

RabbitMQ Connector Initialize

```
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
```

Sample Consumer Definition
```
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
```

Publish Definition
```
event := events.NewSimpleEvent("Hello World!")
rbt.Publish("rabbitmq-test-exchange", "1", event)
```
