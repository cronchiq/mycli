package say

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "say",
	Short: "Say something",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(args[0])
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
