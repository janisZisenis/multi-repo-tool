package version

import (
	"app/log"

	"github.com/spf13/cobra"
)

const CommandName = "version"
const builtBy = "Janis Zisenis"

func MakeCommand(semver string, commit string, date string) *cobra.Command {
	return &cobra.Command{
		Use:   CommandName,
		Short: "Print the version of mrt",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("mrt - version " + semver + ", commit " + commit + ", built at " + date + " by " + builtBy + "\n")
		},
	}
}
