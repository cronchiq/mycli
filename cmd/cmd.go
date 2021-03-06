package cmd

import (
	_ "github.com/cronchiq/mycli/cmd/calc"
	_ "github.com/cronchiq/mycli/cmd/calc/diff"
	_ "github.com/cronchiq/mycli/cmd/calc/divide"
	_ "github.com/cronchiq/mycli/cmd/calc/multiply"
	_ "github.com/cronchiq/mycli/cmd/calc/sum"
	_ "github.com/cronchiq/mycli/cmd/gen_docs"
	_ "github.com/cronchiq/mycli/cmd/hello"

	// _ "github.com/cronchiq/mycli/cmd/ahoj"
	_ "github.com/cronchiq/mycli/cmd/k8s"
	_ "github.com/cronchiq/mycli/cmd/k8s/namespace"
	_ "github.com/cronchiq/mycli/cmd/k8s/nodes"
	_ "github.com/cronchiq/mycli/cmd/k8s/pods"
	"github.com/cronchiq/mycli/cmd/root"
	_ "github.com/cronchiq/mycli/cmd/say"
	_ "github.com/cronchiq/mycli/cmd/say/ahoj"
	_ "github.com/cronchiq/mycli/cmd/say/hello"
	_ "github.com/cronchiq/mycli/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
