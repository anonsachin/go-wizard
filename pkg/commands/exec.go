package commands

import (
	"context"
	"fmt"
	"os"
	cmd "os/exec"
)

func Exec(name string, args ...string) {
	command := cmd.CommandContext(context.Background(), name, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	command.Start()
}

func ExecWithEnv(name string, env map[string]string, args ...string) {
	command := cmd.CommandContext(context.Background(), name, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	for k, v := range env {
		command.Env = append(command.Env, fmt.Sprintf("%s=%s", k, v))
	}

	command.Start()
}
