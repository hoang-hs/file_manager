package services

type UpdateMetricService interface {
	IncreaseRequestCount(path string)
	UpdateStatusCode(path string, statusCode int)
	UpdateLatency(path string, took int64)
}
