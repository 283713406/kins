package cmd

import (
	"io"

	phases "github.com/283713406/kins/app/cmd/phases/apps"
	"github.com/283713406/kins/app/cmd/phases/workflow"
	util "github.com/283713406/kins/app/utils"
	"github.com/spf13/cobra"
)

func NewCmdApps(out io.Writer) *cobra.Command {

	initRunner := workflow.NewRunner()

	cmd := &cobra.Command{
		Use:   "apps",
		Short: "Install apps",
		Run: func(cmd *cobra.Command, args []string) {
			err := initRunner.Run(args)
			util.CheckErr(err)
		},
	}

	initRunner.AppendPhase(phases.NewDeployAppsPhase())

	return cmd
}
