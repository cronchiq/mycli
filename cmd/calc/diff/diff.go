package diff

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/calc"
	"github.com/spf13/cobra"
)



var CmdFlagA int
var CmdFlagB int

var Cmd = &cobra.Command{
	Use:   "diff",
	Short: "Difference of 2 numbers",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Diff: %d - %d = %d \n", CmdFlagA, CmdFlagB, CmdFlagA - CmdFlagB)
	},
}

func init() {
	calc.Cmd.AddCommand(Cmd)

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