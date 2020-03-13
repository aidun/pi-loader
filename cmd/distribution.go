package cmd

import (
	"github.com/spf13/cobra"
)

func MakeDistribution() *cobra.Command {

	var command = &cobra.Command{
		Use:          "distribution",
		Short:        "manage the avaiable distributions",
		Example:      "pi-loader distribution ...",
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		cmd.Help()
	}

	// subcommands
	command.AddCommand(MakeDistributionList())
	return command
}
