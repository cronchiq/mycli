package root

import (
	"github.com/cronchiq/mycli/version"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var RootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My CLI, " + version.Version,
}

func init() {

}