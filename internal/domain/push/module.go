package push

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPushService),
	fx.Provide(NewMemorySubscriptionRepo),
)
