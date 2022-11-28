package cmd

import (
	"fmt"
	"io"

	"github.com/283713406/kins/app/constants"
	"github.com/spf13/cobra"
)

func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of kins",
		Run: func(cmd *cobra.Command, args []string) {
			RunVersion(out, cmd)
		},
	}
	cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}

func RunVersion(out io.Writer, cmd *cobra.Command) error {
	fmt.Printf("kins: %s", constants.Version)
	return nil
}