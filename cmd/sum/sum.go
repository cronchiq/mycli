package sum

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/root"
	"github.com/spf13/cobra"
)



var CmdFlagA int
var CmdFlagB int

var Cmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum of 2 numbers",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(CmdFlagA + CmdFlagB)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)

	Cmd.PersistentFlags().IntVarP(
		&CmdFlagA,
		"x",
		"x",
		0,
		"1st number",
	)
	Cmd.MarkPersistentFlagRequired("x")

	Cmd.PersistentFlags().IntVarP(
		&CmdFlagB,
		"y",
		"y",
		0,
		"2nd number",
	)

	Cmd.MarkPersistentFlagRequired("y")
}