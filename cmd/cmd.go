package cmd

import (
	"github.com/agung96tm/go-learn-proto/cmd/runapi"
	"github.com/agung96tm/go-learn-proto/cmd/runrpc"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(runapi.StartCmd)
	rootCmd.AddCommand(runrpc.StartCmd)
}

var rootCmd = &cobra.Command{
	Use:               "go-miblog-proto",
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run:               func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
