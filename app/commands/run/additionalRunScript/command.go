package additionalRunScript

import (
	"app/core"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const ScriptsPath = "/run/*/command"

func ForScriptInPathDo(path string, do func(scriptPath string, scriptName string)) {
	scripts, _ := filepath.Glob(core.GetExecutablePath() + path)

	for _, script := range scripts {
		segments := strings.Split(script, "/")
		scriptName := segments[len(segments)-2]

		do(script, scriptName)
	}
}

func MakeCommand(scriptName string, scriptPath string) *cobra.Command {
	var command = &cobra.Command{
		Use:   scriptName,
		Short: "Executes additional run script: " + scriptName,
		Run: func(cmd *cobra.Command, args []string) {
			command(scriptPath)
		},
	}

	return command
}

func command(scriptPath string) {
	args := []string{core.GetExecutablePath()}
	script := exec.Command(scriptPath, args...)
	script.Stdout = os.Stdout
	script.Stdin = os.Stdin
	script.Stderr = os.Stderr
	_ = script.Run()
}
