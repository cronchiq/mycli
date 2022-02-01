package calc

import (
	"github.com/cronchiq/mycli/cmd/root"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var Cmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculator",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}