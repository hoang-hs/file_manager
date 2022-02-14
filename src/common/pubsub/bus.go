package pubsub

import "file_manager/src/common/log"

type EventProducer interface {
	ProduceEvent() chan Event
}

type EventBus struct {
	producer     EventProducer
	eventMapping map[string][]Subscriber
}

func NewEventBus(evenProducer EventProducer) *EventBus {
	return &EventBus{
		producer:     evenProducer,
		eventMapping: make(map[string][]Subscriber),
	}
}

func (b *EventBus) Subscribe(event Event, subscriber ...Subscriber) {
	if _, exist := b.eventMapping[event.GetName()]; exist == false {
		b.eventMapping[event.GetName()] = make([]Subscriber, 0)
	}
	b.eventMapping[event.GetName()] = append(b.eventMapping[event.GetName()], subscriber...)
}

func (b *EventBus) Run() {
	for {
		event := <-b.producer.ProduceEvent()
		log.Infof("event [%s] was fired", event.GetName())
		for _, subscriber := range b.eventMapping[event.GetName()] {
			go subscriber.Handle(event)
		}
	}
}
