package divide

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/calc"
	"github.com/spf13/cobra"
)



var CmdFlagA int
var CmdFlagB int

var Cmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide of 2 numbers",	
	Aliases: []string{"dev"},
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if CmdFlagB != 0 {
			fmt.Printf("Divide: %d * %d = %d \n", CmdFlagA, CmdFlagB, CmdFlagA / CmdFlagB)
			
		} else {
			fmt.Println("Mustn't divide by zero")
		}
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