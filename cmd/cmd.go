package cmd

import (
	_ "github.com/cronchiq/mycli/cmd/hello"
	"github.com/cronchiq/mycli/cmd/root"	
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}