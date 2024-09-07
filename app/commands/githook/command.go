package githook

import (
	"app/core"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var CommandName = "git-hook"
var repositoryPath = "repository-path"
var hookNameFlag = "hook-name"

func MakeCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   CommandName,
		Short: "Executes the specified git-hook for a specified repository",
		Run:   command,
	}

	command.Flags().String(hookNameFlag, "", "The name of the git-hook to be executed")
	command.Flags().String(repositoryPath, ".", "The path to the repository")

	return command
}

func command(cmd *cobra.Command, args []string) {
	var teamInfo = core.LoadTeamConfiguration()
	hookName, _ := cmd.Flags().GetString(hookNameFlag)
	repositoryPath, _ := cmd.Flags().GetString(repositoryPath)

	currentBranchName := getCurrentBranchName(repositoryPath)

	switch hookName {
	case core.PreCommit:
		failIfBranchIsBlocked(teamInfo, currentBranchName, "commit")
	case core.PrePush:
		failIfBranchIsBlocked(teamInfo, currentBranchName, "push")
	case core.CommitMsg:
		prefixCommitMessage(teamInfo, currentBranchName, args)
	default:
		fmt.Println("The given git-hook \"" + hookName + "\" does not exist.")
		os.Exit(1)
	}

	executeAdditionalScripts(repositoryPath, hookName, args)
}

func executeAdditionalScripts(repositoryPath string, hookName string, args []string) {
	files, _ := filepath.Glob(repositoryPath + "/hook-scripts/" + hookName + "/*")
	for _, file := range files {
		err := core.ExecuteBash(file, args)

		if err != nil {
			os.Exit(1)
		}
	}
}

func getCurrentBranchName(repositoryPath string) string {
	repository, openError := git.PlainOpen(repositoryPath)

	if openError != nil {
		fmt.Println("The given path \"" + repositoryPath + "\" does not contain a repository.")
		os.Exit(1)
	}

	currentBranch, _ := repository.Head()
	currentBranchName := currentBranch.Name().Short()
	return currentBranchName
}
