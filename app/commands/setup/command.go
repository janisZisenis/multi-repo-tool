package setup

import (
	"app/commands/setup/all"
	"app/commands/setup/cloneRepositories"
	"github.com/spf13/cobra"
)

const commandName = "setup"

func MakeCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   commandName,
		Short: "Sets up you machine for development",
	}

	command.AddCommand(all.MakeCommand())
	command.AddCommand(cloneRepositories.MakeCommand())

	return command
}
