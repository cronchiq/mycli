package say

import (
	"github.com/cronchiq/mycli/cmd/root"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var Cmd = &cobra.Command{
	Use:   "say",
	Short: "Say something",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}