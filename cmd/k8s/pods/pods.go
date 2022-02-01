package pods

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
	Use:     "pods",
	Short:   "Pods list",
	Aliases: []string{"po"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, namespace, _ := k8s.KubernetesClient()

		podsClient := clientset.CoreV1().Pods(namespace)

		fmt.Println("Pods")
		pods, _ := podsClient.List(context.TODO(), metav1.ListOptions{})
		for _, pod := range pods.Items {
			fmt.Println(pod.ObjectMeta.Name, pod.Spec.Containers[0].Image)

		}
	},
}

func init() {
	k8s.Cmd.AddCommand(Cmd)
}
