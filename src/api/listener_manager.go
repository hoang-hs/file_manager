package api

import (
	"file_manager/src/common/pubsub"
	"file_manager/src/core/events"
	"file_manager/src/core/listeners"
)

type ListenersManager struct {
	graphiteListener *listeners.GraphiteListener
}

func NewListenersManager(
	graphiteListener *listeners.GraphiteListener,
	eventBus *pubsub.EventBus,
) *ListenersManager {
	l := &ListenersManager{
		graphiteListener: graphiteListener,
	}
	for event, subscribers := range l.eventMappings() {
		eventBus.Subscribe(event, subscribers...)
	}
	go eventBus.Run()
	return l
}

func (l *ListenersManager) eventMappings() map[pubsub.Event][]pubsub.Subscriber {
	return map[pubsub.Event][]pubsub.Subscriber{
		new(events.RequestEvent): {
			l.graphiteListener,
		},
	}
}
