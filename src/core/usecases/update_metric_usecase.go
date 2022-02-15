package usecases

import (
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/core/entities"
	"file_manager/src/core/enums"
	"fmt"
	"github.com/marpaia/graphite-golang"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type UpdateMetricUseCase struct {
	lockCheckExist *sync.Mutex
	domain         string
	lockMetric     *sync.RWMutex
	metric         map[string]*entities.Metric
	graphite       *graphite.Graphite
}

func NewUpdateMetricUseCase(graphite *graphite.Graphite) services.UpdateMetricService {
	u := &UpdateMetricUseCase{
		lockCheckExist: &sync.Mutex{},
		domain:         "file_manager",
		lockMetric:     &sync.RWMutex{},
		metric:         make(map[string]*entities.Metric, 0),
		graphite:       graphite,
	}
	go u.sendMetric()
	return u
}

func (u *UpdateMetricUseCase) IncreaseRequestCount(path string) {
	u.checkApiMetricExist(path)
	atomic.AddInt64(&u.metric[path].RequestCount, 1)
}

func (u *UpdateMetricUseCase) UpdateStatusCode(path string, statusCode int) {
	u.checkApiMetricExist(path)
	u.lockMetric.Lock()
	u.metric[path].StatusCode[statusCode] += int64(1)
	u.lockMetric.Unlock()
}

func (u *UpdateMetricUseCase) UpdateLatency(path string, took int64) {
	u.checkApiMetricExist(path)
	u.lockMetric.Lock()
	u.metric[path].Latency.LastMinuteRequestTimes = append(u.metric[path].Latency.LastMinuteRequestTimes, float64(took))
	u.lockMetric.Unlock()
}

func (u *UpdateMetricUseCase) checkApiMetricExist(path string) {
	if _, exist := u.metric[path]; exist {
		return
	}
	u.lockCheckExist.Lock()
	if _, exist := u.metric[path]; exist {
		return
	}
	u.metric[path] = entities.NewMetric()
	u.lockCheckExist.Unlock()
}

func (u *UpdateMetricUseCase) sendMetric() {
	for range time.NewTicker(enums.TimeUpdate).C {
		metrics := u.generateGraphiteMetrics()
		err := u.graphite.SendMetrics(metrics)
		if err != nil {
			log.Errorf("Can not send metric to graphite, error: [%s]", err)
		} else {
			log.Infof("Sent metrics to graphite successfully")
		}
	}
}

func (u *UpdateMetricUseCase) generateGraphiteMetrics() []graphite.Metric {
	timestamp := time.Now().Unix()
	metrics := make([]graphite.Metric, 0)
	for api, metric := range u.metric {
		nameMetric := fmt.Sprintf("%s.%s.%s", u.domain, api, "request")
		v := metric.RequestCount
		metrics = append(metrics, graphite.NewMetric(nameMetric, strconv.FormatInt(v, 10), timestamp))

		for statusCode, v := range metric.StatusCode {
			nameMetric = fmt.Sprintf("%s.%s.%s.%v", u.domain, api, "status_code", statusCode)
			metrics = append(metrics, graphite.NewMetric(nameMetric, strconv.FormatInt(v, 10), timestamp))
		}

		metric.Latency.Calculate()
		nameMetric = fmt.Sprintf("%s.%s.%s", u.domain, api, "latency.min")
		vLatency := int(metric.Latency.Min)
		metrics = append(metrics, graphite.NewMetric(nameMetric, strconv.Itoa(vLatency), timestamp))

		nameMetric = fmt.Sprintf("%s.%s.%s", u.domain, api, "latency.max")
		vLatency = int(metric.Latency.Max)
		metrics = append(metrics, graphite.NewMetric(nameMetric, strconv.Itoa(vLatency), timestamp))

		nameMetric = fmt.Sprintf("%s.%s.%s", u.domain, api, "latency.mean")
		vLatency = int(metric.Latency.Mean)
		metrics = append(metrics, graphite.NewMetric(nameMetric, strconv.Itoa(vLatency), timestamp))

		metric.ResetMetric()
	}
	return metrics
}
