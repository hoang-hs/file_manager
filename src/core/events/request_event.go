package events

type RequestMessage struct {
	Path        string
	StatusCode  int
	LatencyData int64
}

type RequestEvent struct {
	PayloadData *RequestMessage
}

func NewRequestMessage(path string, took int64) *RequestEvent {
	return &RequestEvent{
		PayloadData: &RequestMessage{
			Path:        path,
			LatencyData: took,
		},
	}
}

func (r *RequestEvent) GetName() string {
	return "request_event"
}

func (r *RequestEvent) Payload() interface{} {
	return r.PayloadData
}
