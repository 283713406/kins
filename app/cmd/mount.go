package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func NewCmdMount(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mount",
		Short: "Mount glusterfs",
		Run: func(cmd *cobra.Command, args []string) {
			RunMount(out, cmd)
		},
	}
	cmd.Flags().StringP("output", "o", "", "Output format; available options are 'yaml', 'json' and 'short'")
	return cmd
}

func RunMount(out io.Writer, cmd *cobra.Command) error {
	fmt.Printf("[Mount] Mount glusterfs")

	return nil
}