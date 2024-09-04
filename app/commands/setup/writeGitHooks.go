package setup

import (
	"app/core"
	"fmt"
	"os"
)

func getHookTemplate() string {
	return `
#!/bin/bash -e

hook_name=$(basename "$0")

branch="$(git rev-parse --abbrev-ref HEAD)"
` + core.GetExecutable() + ` githook --hook-name "$hook_name" --branch $branch $@`
}

func writeGitHook(repositoryDirectory string, hookName string) {
	hooksPath := repositoryDirectory + "/.git/hooks/"
	_ = os.MkdirAll(hooksPath, os.ModePerm)
	err := os.WriteFile(hooksPath+hookName, []byte(getHookTemplate()), 0755)
	if err != nil {
		fmt.Printf("unable to write file: %w", err)
	}
}