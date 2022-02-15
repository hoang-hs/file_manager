package entities

import (
	"sort"
	"sync"
)

type Latency struct {
	lock                   *sync.RWMutex
	LastMinuteRequestTimes []float64
	Min                    float64 `json:"min"`
	Max                    float64 `json:"max"`
	Mean                   float64 `json:"mean"`
}

func NewLatency() *Latency {
	return &Latency{
		LastMinuteRequestTimes: make([]float64, 0),
		lock:                   &sync.RWMutex{},
		Min:                    0,
		Max:                    0,
		Mean:                   0,
	}
}

func (l *Latency) Reset() {
	l.LastMinuteRequestTimes = make([]float64, 0)
	l.Min = 0
	l.Max = 0
	l.Mean = 0
}

func (l *Latency) Calculate() {
	l.lock.Lock()
	defer l.lock.Unlock()
	sortedSlice := l.LastMinuteRequestTimes[:]
	l.LastMinuteRequestTimes = make([]float64, 0)
	length := len(sortedSlice)
	if length < 1 {
		return
	}
	sort.Float64s(sortedSlice)
	l.Min = sortedSlice[0]
	l.Max = sortedSlice[length-1]
	l.Mean = l.mean(sortedSlice, length)
}

func (l *Latency) mean(orderedObservations []float64, length int) float64 {
	res := 0.0
	for i := 0; i < length; i++ {
		res += orderedObservations[i]
	}

	return res / float64(length)
}
