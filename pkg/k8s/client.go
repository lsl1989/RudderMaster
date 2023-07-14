package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Clients struct {
	ClientSet *kubernetes.Clientset
}

func NewClient(confPath string) (clients Clients, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", confPath)
	if err != nil {
		return
	}
	clients.ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	return
}
