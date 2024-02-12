package bootstrap

import (
	"context"
	"fmt"
	"github.com/agung96tm/go-learn-proto/apps/api/controllers"
	"github.com/agung96tm/go-learn-proto/apps/api/routes"
	"github.com/agung96tm/go-learn-proto/lib"
	"go.uber.org/fx"
)

var ApiModule = fx.Module(
	"api",
	fx.Provide(lib.NewHttpHandler),
	controllers.Module,
	routes.Module,
	fx.Invoke(apiBootstrap),
)

func apiBootstrap(lifecycle fx.Lifecycle, handler lib.HttpHandler, routes routes.Routes) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			routes.Setup()

			go func() {
				if err := handler.Run(); err != nil {
					panic(err)
				}
				fmt.Println("API Bootstrap Running....")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("API Bootstrap Stop....")
			return nil
		},
	})
}
