package cmd

import (
	"os"
	"os/signal"
	"time"

	"github.com/balchua/pod-finder/handler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientSet kubernetes.Clientset
var done chan bool
var ticker *time.Ticker

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "pod-finder is a helper command to get pod status.",
	Long:  "A simple go application which finds all the pods based on a label and prints to a file the pod names and IP addresses.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ticker = time.NewTicker(time.Duration(Period) * time.Second)
		c := make(chan os.Signal)
		done = make(chan bool, 1)
		signal.Notify(c, os.Interrupt)
		log.Infof("checking pods in namespace [%s]", Namespace)
		kclient, err := getClient()
		if err != nil {
			log.Warnf("Failed to get kubernetes client: %v", err)
			return err
		}

		go signalTerminate(c)

		go doWork(kclient)

		<-done

		log.Info("One last run before shutting down.")
		findPods(kclient)
		return nil
	},
}

func init() {

	rootCmd.AddCommand(checkCmd)

}

func getClient() (*kubernetes.Clientset, error) {
	log.Debug("Calling getClient()")
	var config *rest.Config
	var err error

	if PathToConfig == "" {
		log.Info("Using in cluster config")
		config, err = rest.InClusterConfig()
	} else {
		log.Info("Using out of cluster config")
		config, err = clientcmd.BuildConfigFromFlags("", PathToConfig)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func getPods(kclient *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {

	podClient := kclient.CoreV1().Pods(namespace)

	var listOptions metav1.ListOptions

	if Selectors != "" {
		listOptions = metav1.ListOptions{
			LabelSelector: Selectors,
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

func doWork(kclient *kubernetes.Clientset) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			findPods(kclient)
		}
	}
}

func findPods(kclient *kubernetes.Clientset) {
	pods, err := getPods(kclient, Namespace)
	if err != nil {
		log.Error("Unable to get pods.")
	}
	handler.WriteToFile(pods)
}

func signalTerminate(c <-chan os.Signal) {
	select {
	case sig := <-c:
		log.Info("Got %s signal. Aborting...\n", sig)
		stop()
	}
}

func stop() {
	ticker.Stop()
	close(done)
}
