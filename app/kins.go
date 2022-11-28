package app

import (
	"os"

	"github.com/283713406/kins/app/cmd"
)

// Run creates and executes new kins command
func Run() error {
	cmd := cmd.NewKinsCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
