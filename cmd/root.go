package cmd

import (
	"os"

	"github.com/kujourinka/mhws_beta_server/backend"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "...",
	Run: mainRun,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func mainRun(cmd *cobra.Command, args []string) {
	e := backend.RegisterHandler()

	e.RunTLS(":443", "cert/website.crt", "cert/website.key")
}
