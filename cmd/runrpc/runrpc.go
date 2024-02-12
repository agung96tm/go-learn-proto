package runrpc

import (
	"github.com/agung96tm/go-learn-proto/bootstrap"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var StartCmd = &cobra.Command{
	Use:   "runrpc",
	Short: "Start Application as RPC",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	fx.New(bootstrap.RpcModule).Run()
}
