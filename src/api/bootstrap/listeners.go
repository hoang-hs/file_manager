package bootstrap

import (
	"file_manager/src/api/listeners"
	"file_manager/src/common/pubsub"
	"go.uber.org/fx"
)

func LoadListeners() []fx.Option {
	return []fx.Option{
		fx.Provide(pubsub.NewPublisher),
		fx.Invoke(pubsub.RegisterGlobal),
		fx.Provide(pubsub.NewEventBus),
		fx.Provide(listeners.NewUpdateRequestCountListener),
		fx.Provide(listeners.NewUpdateStatusCodeResponseListener),
		fx.Provide(listeners.NewListenersManager),
		fx.Invoke(listeners.Subscribes),
	}
}
