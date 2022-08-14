package httphandlers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewTodoHandler,
	),
	fx.Invoke(
		ConfigureHealthCheckHandler,
		ConfigureTodoHandler,
	),
)
