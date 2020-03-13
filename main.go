package main

import (
	"os"

	"github.com/aidun/pi-loader/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use: "pi-loader",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Short:        "create and flash images to sd-cards",
		Example:      "pi-loader ...",
		SilenceUsage: false,
	}

	rootCmd.AddCommand(cmd.MakeFlash())

	//rootCmd.GenBashCompletion(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
