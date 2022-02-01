package multiply

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/calc"
	"github.com/spf13/cobra"
)



var CmdFlagA int
var CmdFlagB int

var Cmd = &cobra.Command{
	Use:   "multiply",
	Short: "Multiply of 2 numbers",	
	Aliases: []string{"mul"},
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Multiply: %d * %d = %d \n", CmdFlagA, CmdFlagB, CmdFlagA * CmdFlagB)
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