package cmd

import (
	"io"

	phases "github.com/283713406/kins/app/cmd/phases/basic"
	"github.com/283713406/kins/app/cmd/phases/workflow"
	util "github.com/283713406/kins/app/utils"
	"github.com/spf13/cobra"
)

func NewCmdBasic(out io.Writer) *cobra.Command {

	initRunner := workflow.NewRunner()

	cmd := &cobra.Command{
		Use:   "basic",
		Short: "Install basic",
		Run: func(cmd *cobra.Command, args []string) {
			err := initRunner.Run(args)
			util.CheckErr(err)
		},
	}

	initRunner.AppendPhase(phases.NewDeployBasicPhase())

	return cmd
}