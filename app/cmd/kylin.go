package cmd

import (
	"fmt"
	"io"

	"github.com/283713406/kins/app/constants"
	"github.com/spf13/cobra"
)

func NewCmdKylin(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kylin",
		Short: "Install kylin platform",
		Run: func(cmd *cobra.Command, args []string) {
			RunKylin(out, cmd)
		},
	}
	cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}

func RunKylin(out io.Writer, cmd *cobra.Command) error {
	fmt.Printf("kins: %s", constants.Version)
	return nil
}