package conf

import "go.uber.org/fx"

var Module = fx.Module("appconfigfx",
	// - order is not important in provide
	// - provide can have parameter and will resolve if registered
	// - execute its func only if it requested
	fx.Provide(
		NewAppOptions,
	),
)
