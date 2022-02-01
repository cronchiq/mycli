package cmd

import (
	_ "github.com/cronchiq/mycli/cmd/hello"
	_ "github.com/cronchiq/mycli/cmd/calc"
	_ "github.com/cronchiq/mycli/cmd/calc/sum"
	_ "github.com/cronchiq/mycli/cmd/calc/diff"
	// _ "github.com/cronchiq/mycli/cmd/ahoj"
	_ "github.com/cronchiq/mycli/cmd/sum"
	_ "github.com/cronchiq/mycli/cmd/say"
	_ "github.com/cronchiq/mycli/cmd/say/ahoj"
	_ "github.com/cronchiq/mycli/cmd/say/hello"
	"github.com/cronchiq/mycli/cmd/root"	
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}