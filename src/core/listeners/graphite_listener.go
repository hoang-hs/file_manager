package listeners

import (
	"file_manager/src/common/log"
	"file_manager/src/common/pubsub"
	"file_manager/src/core/events"
	"file_manager/src/core/usecases"
)

type GraphiteListener struct {
	updateMetricUseCase *usecases.UpdateMetricUseCase
}

func NewGraphiteListener(
	updateMetricUseCase *usecases.UpdateMetricUseCase,
) *GraphiteListener {
	return &GraphiteListener{
		updateMetricUseCase: updateMetricUseCase,
	}
}

func (g *GraphiteListener) Handle(event pubsub.Event) {
	e := event.(*events.RequestEvent)
	log.Infof("[GraphiteListener] received an event: [%s] ", e.GetName())
	g.updateMetricUseCase.IncreaseRequestCount(e.PayloadData.Path)
	g.updateMetricUseCase.UpdateLatency(e.PayloadData.Path, e.PayloadData.LatencyData)
	g.updateMetricUseCase.UpdateStatusCode(e.PayloadData.Path, e.PayloadData.StatusCode)
}
