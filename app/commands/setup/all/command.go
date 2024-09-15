package all

import (
	"app/commands/setup/additionalScript"
	"app/commands/setup/cloneRepositories"
	"app/commands/setup/installGitHooks"
	"app/log"
	"github.com/spf13/cobra"
)

const skipHooksFlag = "skip-git-hooks"
const scriptName = "all"

func MakeCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   scriptName,
		Short: "Executes all setup commands",
		Run:   command,
	}

	command.Flags().Bool(skipHooksFlag, false, "Skips setting the git-hooks")
	command.Flags().Lookup(skipHooksFlag).NoOptDefVal = "true"

	return command
}

func command(cmd *cobra.Command, args []string) {
	shouldSkipHooks, _ := cmd.Flags().GetBool(skipHooksFlag)

	cloneRepositories.MakeCommand().Run(cmd, args)

	if !shouldSkipHooks {
		installGitHooks.MakeCommand().Run(cmd, args)
	}

	executeAdditionalSetupScripts(cmd, args)
}

func executeAdditionalSetupScripts(cmd *cobra.Command, args []string) {
	log.Info("Executing additional setup-scripts.")

	additionalScript.ForScriptInPathDo(additionalScript.ScriptsPath, func(scriptPath string) {
		additionalScript.MakeCommand(scriptPath).Run(cmd, args)
	})

	log.Success("Done executing additional setup-scripts.")
}
