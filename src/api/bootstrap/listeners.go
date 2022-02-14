package bootstrap

import (
	"file_manager/src/api"
	"file_manager/src/common/pubsub"
	"file_manager/src/core/listeners"
	"go.uber.org/fx"
)

func LoadListeners() []fx.Option {
	return []fx.Option{
		fx.Provide(pubsub.NewPublisher),
		fx.Invoke(pubsub.RegisterGlobal),
		fx.Provide(pubsub.NewEventBus),
		fx.Provide(listeners.NewGraphiteListener),
		fx.Provide(api.NewListenersManager),
	}
}
