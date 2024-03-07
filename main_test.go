package main_test

import (
	"os"
	"testing"

	"github.com/nlnwa/warchaeology/cmd"
)

func TestEmptyCommandPrompt(t *testing.T) {
	stdout_backup := os.Stdout
	stderr_backup := os.Stderr
	os.Stdout = nil
	os.Stderr = nil

	shell := cmd.NewCommand()
	_ = shell.Execute()

	os.Stdout = stdout_backup
	os.Stderr = stderr_backup
}
