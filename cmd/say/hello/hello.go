package hello

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/say"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "Say: Hello!",
	Aliases: []string{"h"},
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}

func init() {
	say.Cmd.AddCommand(Cmd)
}