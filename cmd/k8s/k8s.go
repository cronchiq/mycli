package k8s

import (
	"github.com/cronchiq/mycli/cmd/root"
	"github.com/spf13/cobra"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func KubernetesClient() (*kubernetes.Clientset, string, error) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	namespace, _, _ := kubeConfig.Namespace()
	restconfig, _ := kubeConfig.ClientConfig()

	clientset, _ := kubernetes.NewForConfig(restconfig)
	return clientset, namespace, nil
}

var Cmd = &cobra.Command{
	Use:   "k8s",
	Short: "k8s API!",
	/* Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, namespace, _ := KubernetesClient()

		// Namespace
		fmt.Println("Namespace: " + namespace)

		nodesClient := clientset.CoreV1().Nodes()
		podsClient := clientset.CoreV1().Pods(namespace)

		// Nodes
		fmt.Println("Nodes:")

		nodes, _ := nodesClient.List(context.TODO(), metav1.ListOptions{})
		for _, node := range nodes.Items {
			fmt.Println(node.ObjectMeta.Name, node.Status.Addresses[0].Address)

		}

		// Pods
		fmt.Println("Pods")
		pods, _ := podsClient.List(context.TODO(), metav1.ListOptions{})
		for _, pod := range pods.Items {
			fmt.Println(pod.ObjectMeta.Name, pod.Spec.Containers[0].Image)

		}

	}, */
}

func init() {
	root.RootCmd.AddCommand(Cmd)

}
