package version

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/root"
	"github.com/cronchiq/mycli/version"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "My CLI version",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}