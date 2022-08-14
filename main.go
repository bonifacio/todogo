package main

import (
	"context"
	"log"

	"github.com/bonifacio/todogo/handlers/httphandlers"
	"github.com/bonifacio/todogo/repositories"
	"github.com/bonifacio/todogo/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	fx.New(
		fx.Provide(gin.Default),
		services.Module,
		httphandlers.Module,
		repositories.Module,
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lifeCycle fx.Lifecycle, engine *gin.Engine) {
	lifeCycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go engine.Run(":8080")
				return nil
			},
		},
	)
}
