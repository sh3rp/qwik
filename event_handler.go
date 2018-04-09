package qwik

import "fmt"

type EventHandler struct {
	publisher   EventPublisher
	transformer EventTransformer
}

func NewEventHandler(publisher EventPublisher, transformer EventTransformer) EventHandler {
	return EventHandler{
		publisher:   publisher,
		transformer: transformer,
	}
}

func (eh EventHandler) Handle(evt FSEvent) {
	fmt.Printf("Event: %v\n", evt)
	eh.publisher.Publish(eh.transformer.FromFSEvent(evt))
}
