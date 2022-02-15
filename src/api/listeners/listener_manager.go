package listeners

import (
	"file_manager/src/common/pubsub"
	"file_manager/src/core/events"
)

type ListenerManager struct {
	updateRequestCountListener       *UpdateRequestCountListener
	updateStatusCodeResponseListener *UpdateStatusCodeResponseListener
}

func NewListenersManager(
	updateRequestCountListener *UpdateRequestCountListener,
	updateStatusCodeResponseListener *UpdateStatusCodeResponseListener,
) *ListenerManager {
	return &ListenerManager{
		updateRequestCountListener:       updateRequestCountListener,
		updateStatusCodeResponseListener: updateStatusCodeResponseListener,
	}
}

func (l *ListenerManager) eventMappings() map[pubsub.Event][]pubsub.Subscriber {
	return map[pubsub.Event][]pubsub.Subscriber{
		new(events.RequestEvent): {
			l.updateRequestCountListener,
		},
		new(events.ResponseEvent): {
			l.updateStatusCodeResponseListener,
		},
	}
}

func Subscribes(l *ListenerManager, eventBus *pubsub.EventBus) {
	for event, subscribers := range l.eventMappings() {
		eventBus.Subscribe(event, subscribers...)
	}
	go eventBus.Run()
}
