package runapi

import (
	"github.com/agung96tm/go-learn-proto/bootstrap"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var StartCmd = &cobra.Command{
	Use:   "runapi",
	Short: "Start Application as API",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func runApplication() {
	fx.New(bootstrap.ApiModule).Run()
}
