package main

import (
	"api-bed-covid/cli/commands"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "abc",
		Short: "Example CLI command in Go using cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCmd.AddCommand(
		commands.ListProvinces(),
		commands.CheckSupabase(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
