package bootstrap

import (
	"context"
	"fmt"
	appRpc "github.com/agung96tm/go-learn-proto/apps/rpc"
	"github.com/agung96tm/go-learn-proto/lib"
	"go.uber.org/fx"
)

var RpcModule = fx.Options(
	fx.Provide(lib.NewRpcHandler),
	appRpc.Module,
	fx.Invoke(rpcBootstrap),
)

func rpcBootstrap(lifecycle fx.Lifecycle, services appRpc.Services, handler lib.RpcHandler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				services.Setup()

				if err := handler.Run(); err != nil {
					panic(err)
				}

				fmt.Println("RPC Application Start....")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("RPC Application Stop....")
			return nil
		},
	})
}
