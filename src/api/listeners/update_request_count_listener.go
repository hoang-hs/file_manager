package listeners

import (
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/common/pubsub"
	"file_manager/src/core/events"
)

type UpdateRequestCountListener struct {
	updateMetricService services.UpdateMetricService
}

func NewUpdateRequestCountListener(
	updateMetricService services.UpdateMetricService,
) *UpdateRequestCountListener {
	return &UpdateRequestCountListener{
		updateMetricService: updateMetricService,
	}
}

func (g *UpdateRequestCountListener) Handle(event pubsub.Event) {
	e, ok := event.(*events.RequestEvent)
	if !ok {
		log.Errorf("UpdateStatusCodeResponseListener received a invalid event: %v", event)
		return
	}
	log.Infof("[UpdateRequestCountListener] received an event: [%s] ", e.GetName())
	g.updateMetricService.IncreaseRequestCount(e.PayloadData.Path)
	g.updateMetricService.UpdateLatency(e.PayloadData.Path, e.PayloadData.LatencyData)
}
