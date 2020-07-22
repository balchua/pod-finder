package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientSet kubernetes.Clientset
var selectors string
var pathToConfig string

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("A walrus appears")
	flag.StringVar(&pathToConfig, "kubeConfig", "", "Text to parse.")

	flag.Parse()
	pods, _ := getPods("artemis")

	for _, pod := range pods.Items {
		log.Infof("Pod name[%s], Pod Ip [%s],Pod status [%s]", pod.Name, pod.Status.PodIP, pod.Status.Phase)
	}
}

func getClient() (*kubernetes.Clientset, error) {
	log.Debug("Calling getClient()")
	var config *rest.Config
	var err error

	if pathToConfig == "" {
		log.Info("Using in cluster config")
		config, err = rest.InClusterConfig()
	} else {
		log.Info("Using out of cluster config")
		config, err = clientcmd.BuildConfigFromFlags("", pathToConfig)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func getPods(namespace string) (*corev1.PodList, error) {
	kclient, err := getClient()
	if err != nil {
		log.Warnf("Failed to get kubernetes client: %v", err)
		return nil, err
	}
	podClient := kclient.CoreV1().Pods(namespace)

	var listOptions metav1.ListOptions

	if selectors != "" {
		listOptions = metav1.ListOptions{
			LabelSelector: selectors,
			Limit:         100,
		}

	} else {
		listOptions = metav1.ListOptions{}
	}
	podList, err := podClient.List(listOptions)

	if err != nil {
		log.Warnf("Failed to find pods: %v", err)
		return nil, err
	}

	return podList, nil
}
