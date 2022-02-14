package pubsub

type Event interface {
	GetName() string
	Payload() interface{}
}
