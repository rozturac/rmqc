# 🐇 RMQC (RabbitMQ Connector)

[![GitHub license](https://img.shields.io/github/license/rozturac/rmqc.svg)](https://github.com/rozturac/rmqc/LICENSE)

RMQC provide to work with RabbitMQ without complexity. RMQC uses AMQP framework in background and does not reflect the complexity that comes with amqp to the developer.

## Installation

Via go packages:
```go get github.com/rozturac/rmqc```

## Usage

### Configure

*Hint:*
Check the [go-ddd-example](https://github.com/rozturac/go-ddd-example) project, the project have rmqc integration with consumers and publisher.

Here is a sample RabbitMQ initialization:

```go
import (
	"github.com/rozturac/rmqc"
	consumers "github.com/rozturac/rmqc/_examples/consumers"
	"github.com/rozturac/rmqc/_examples/events"
)

func NewRabbitMQ() (*rmqc.RabbitMQ, error) {
	return rmqc.Connect(rmqc.RabbitMQConfig{
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
}
```

### Register Event

The Event must be implemented from the IEvent interface.

```go
//Implemented by rmqc.IEvent 
type SimpleEvent struct {
	Message string
}

func NewSimpleEvent(message string) *SimpleEvent {
	return &SimpleEvent{Message: message}
}
```

### Register Event Consumer

Queue and Exchange definitions to be made, consumers must first be configured.

```go
//Implemented by rmqc.IConsumer
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
```

### Publish Event 

```go
rbt := NewRabbitMQ()
event := events.NewSimpleEvent("Hello World!")
rbt.Publish("rabbitmq-test-exchange", "1", event)
```

### Bind Consumers

```go
// Event data structure
rbt := NewRabbitMQ()
consumer := consumers.NewSimpleConsumer()
if err = rbt.BindConsumer(consumer); err != nil {
	panic(err)
}
rbt.Start()
```

### Sample Project

[go-ddd-example](https://github.com/rozturac/go-ddd-example)

## License
MIT License

Copyright (c) 2021 Rıdvan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
