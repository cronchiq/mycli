package namespace

import (
	"fmt"

	"github.com/cronchiq/mycli/cmd/k8s"
	"github.com/spf13/cobra"
)

// func KubernetesClient() (*kubernetes.Clientset, string, error) {
// 	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
// 		clientcmd.NewDefaultClientConfigLoadingRules(),
// 		&clientcmd.ConfigOverrides{},
// 	)
// 	namespace, _, _ := kubeConfig.Namespace()
// 	restconfig, _ := kubeConfig.ClientConfig()

// 	clientset, _ := kubernetes.NewForConfig(restconfig)
// 	return clientset, namespace, nil
// }

var Cmd = &cobra.Command{
	Use:     "namespace",
	Short:   "Get current namespace",
	Aliases: []string{"ns"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		_, namespace, _ := k8s.KubernetesClient()

		// Namespace
		fmt.Println("Namespace: " + namespace)
	},
}

func init() {
	k8s.Cmd.AddCommand(Cmd)
}
