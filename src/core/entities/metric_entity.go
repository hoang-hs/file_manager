package entities

type Metric struct {
	RequestCount int64
	StatusCode   map[int]int64
	Latency      *Latency
}

func NewMetric() *Metric {
	return &Metric{
		StatusCode: make(map[int]int64, 0),
		Latency:    NewLatency(),
	}
}

func (m *Metric) ResetMetric() {
	m.RequestCount = 0
	m.StatusCode = make(map[int]int64, 0)
	m.Latency.Reset()
}
