package cmd

import (
	"github.com/cronchiq/mycli/cmd/root"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}