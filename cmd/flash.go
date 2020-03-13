package cmd

import (
	"github.com/aidun/pi-loader/pkg/repo"
	"github.com/spf13/cobra"
)

func MakeFlash() *cobra.Command {

	var command = &cobra.Command{
		Use:          "flash",
		Short:        "flash an image to sd card",
		Example:      "pi-loader flash ...",
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		repo.Initialize()
	}
	return command
}
