package k8s

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tonygilkerson/ispy/internal/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClientSet() *kubernetes.Clientset {

	inCluster, exists := os.LookupEnv("ISPY_IN_CLUSTER")
	if exists {
		log.Printf("Using environment variable ISPY_IN_CLUSTER: %v", inCluster)
	} else {
		inCluster = "false"
		log.Printf("ISPY_IN_CLUSTER environment variable not set, using default value: %v", inCluster)
	}

	var clientset *kubernetes.Clientset
	var kubeConfig *rest.Config
	var err error

	if inCluster == "true" {
		log.Println("Creating an in cluster go-client config")
		kubeConfig, err = rest.InClusterConfig()
		
		if err != nil {
			// panic(err.Error())
			log.Printf("NOT GOOD: %v\n", err.Error())
		}		
		
	} else {
		log.Panicln("Creating go-client config from user homedir")
		userHomeDir, err := os.UserHomeDir()
		util.DoOrDie(err)

		kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
		fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

		kubeConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
		util.DoOrDie(err)
	}
	
	clientset, err = kubernetes.NewForConfig(kubeConfig)
	util.DoOrDie(err)

	return clientset
}

func GetPods(namespace string, client kubernetes.Interface) (*v1.PodList, error) {
	fmt.Println("Get Kubernetes Pods")
	pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	util.DoOrDie(err)

	return pods, nil
}

func GetNamespaces(client kubernetes.Interface) (*v1.NamespaceList, error) {
	fmt.Println("Get Kubernetes Namespaces")
	namespaces, err := client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	util.DoOrDie(err)

	return namespaces, nil
}
