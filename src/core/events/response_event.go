package events

type ResponseMessage struct {
	Path       string
	StatusCode int
}

type ResponseEvent struct {
	PayloadData *ResponseMessage
}

func NewResponseEvent(path string, statusCode int) *ResponseEvent {
	return &ResponseEvent{
		PayloadData: &ResponseMessage{
			Path:       path,
			StatusCode: statusCode,
		},
	}
}

func (r *ResponseEvent) GetName() string {
	return "response_event"
}

func (r *ResponseEvent) Payload() interface{} {
	return r.PayloadData
}
