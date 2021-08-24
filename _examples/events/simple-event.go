package events

type SimpleEvent struct {
	Message string
}

func NewSimpleEvent(message string) *SimpleEvent {
	return &SimpleEvent{Message: message}
}
