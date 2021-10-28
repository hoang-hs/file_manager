package resources

type Message struct {
	Message string `json:"message"`
}

func NewMessageResource(message string) *Message {
	return &Message{
		Message: message,
	}
}
