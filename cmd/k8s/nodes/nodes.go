package nodes

import (
	"context"
	"fmt"

	"github.com/cronchiq/mycli/cmd/k8s"
	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Use:     "nodes",
	Short:   "Nodes list",
	Aliases: []string{"no"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, _ := k8s.KubernetesClient()

		nodesClient := clientset.CoreV1().Nodes()

		fmt.Println("Nodes:")

		nodes, _ := nodesClient.List(context.TODO(), metav1.ListOptions{})
		for _, node := range nodes.Items {
			fmt.Println(node.ObjectMeta.Name, node.Status.Addresses[0].Address)

		}
	},
}

func init() {
	k8s.Cmd.AddCommand(Cmd)
}
