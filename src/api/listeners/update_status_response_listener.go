package listeners

import (
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/common/pubsub"
	"file_manager/src/core/events"
)

type UpdateStatusCodeResponseListener struct {
	updateMetricService services.UpdateMetricService
}

func NewUpdateStatusCodeResponseListener(
	updateMetricService services.UpdateMetricService,
) *UpdateStatusCodeResponseListener {
	return &UpdateStatusCodeResponseListener{
		updateMetricService: updateMetricService,
	}
}

func (g *UpdateStatusCodeResponseListener) Handle(event pubsub.Event) {
	e, ok := event.(*events.ResponseEvent)
	if !ok {
		log.Errorf("UpdateStatusCodeResponseListener received a invalid event: %v", event.GetName())
		return
	}
	log.Infof("[UpdateStatusCodeResponseListener] received an event: [%s] ", e.GetName())
	g.updateMetricService.UpdateStatusCode(e.PayloadData.Path, e.PayloadData.StatusCode)
}
