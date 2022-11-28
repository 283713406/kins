package cmd

import (
	"io"

	phases "github.com/283713406/kins/app/cmd/phases/k8s"
	"github.com/283713406/kins/app/cmd/phases/workflow"
	util "github.com/283713406/kins/app/utils"
	"github.com/spf13/cobra"
)

func NewCmdK8s(out io.Writer) *cobra.Command {

	initRunner := workflow.NewRunner()

	cmd := &cobra.Command{
		Use:   "k8s",
		Short: "Install kylin container cloud",
		Run: func(cmd *cobra.Command, args []string) {
			err := initRunner.Run(args)
			util.CheckErr(err)
		},
	}

	initRunner.AppendPhase(phases.NewChecksPhase())
	initRunner.AppendPhase(phases.NewMountK8sIsoPhase())
	initRunner.AppendPhase(phases.NewPreparePhase())
	initRunner.AppendPhase(phases.NewLoadK8sImagesPhase())
	initRunner.AppendPhase(phases.NewDeployK8sPhase())

	initRunner.BindToCommand(cmd)

	return cmd
}