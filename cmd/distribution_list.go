package cmd

import (
	"github.com/aidun/pi-loader/pkg/repo"
	"github.com/spf13/cobra"
)

func MakeDistributionList() *cobra.Command {

	var command = &cobra.Command{
		Use:          "list",
		Short:        "list the avaiable distributions",
		Example:      "pi-loader distribution list",
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		repo.ListDistributions()
	}
	return command
}
