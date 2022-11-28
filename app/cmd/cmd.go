package cmd

import (
	"io"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
)

func NewKinsCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	cmds := &cobra.Command{
		Use:   "kins",
		Short: "麒麟终端系统管控平台部署工具",
		Long: dedent.Dedent(`
				
				KINS  
				麒麟终端系统管控平台部署工具，支持一键部署及各组件单独部署	    

			Example usage:

				一键部署
				control-plane# kins kylin install

				单独部署
				worker# kins k8s install

			    You can then repeat the second step on as many other machines as you like.

		`),
	}

	cmds.AddCommand(NewCmdKylin(out))
	cmds.AddCommand(NewCmdK8s(out))
	cmds.AddCommand(NewCmdBasic(out))
	cmds.AddCommand(NewCmdApps(out))
	cmds.AddCommand(NewCmdMount(out))
	cmds.AddCommand(NewCmdVersion(out))

	return cmds
}