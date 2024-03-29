package empty

import (
	"github.com/danilluk1/social-network/libs/go/pkg/logger"
	"github.com/danilluk1/social-network/libs/go/pkg/logger/conf"
	"go.uber.org/fx"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module("emptyfx",
	// - order is not important in provide
	// - provide can have parameter and will resolve if registered
	// - execute its func only if it requested
	fx.Provide(
		fx.Annotate(
			EmptyLogger,
			fx.As(new(logger.Logger)),
		),
		conf.ProvideLogConfig,
	))
