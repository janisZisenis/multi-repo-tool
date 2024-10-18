package core

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ForScriptInPathDo(path string, do func(scriptPath string, scriptName string)) {
	scripts, _ := filepath.Glob(path)

	for _, script := range scripts {
		dirPath := filepath.Dir(script)
		scriptName := filepath.Base(dirPath)

		do(script, scriptName)
	}
}

type ExitCode = int

func ExecuteScript(scriptPath string, args []string) ExitCode {
	script := exec.Command(scriptPath, args...)
	script.Stdout = os.Stdout
	script.Stdin = os.Stdin
	script.Stderr = os.Stderr
	_ = script.Run()

	return 0
}
