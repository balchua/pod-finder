package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Selectors    string
	PathToConfig string
	Namespace    string
	Period       int
	OutputPath   string

	rootCmd = &cobra.Command{

		Use:   "pod-finder",
		Short: "pod-finder is a helper command to get pod status.",
		Long:  "A simple go application which finds all the pods based on a label and prints to a file the pod names and IP addresses.",
	}
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	rootCmd.PersistentFlags().StringVarP(&PathToConfig, "config", "c", "", "config file where the $KUBECONFIG is located.  If empty, it will assume that the pod-finder is running inside the cluster.")
	rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "The namespace to look for the pods.")
	rootCmd.PersistentFlags().StringVarP(&Selectors, "label", "l", "", "To find pods based on kubernetes selector")
	rootCmd.PersistentFlags().IntVarP(&Period, "period", "p", 10, "The period in seconds to get Pod information.")
	rootCmd.PersistentFlags().StringVarP(&OutputPath, "output", "o", "", "(Required) Output file from where the result will be written.  Ex: /tmp/output.json")
	rootCmd.MarkPersistentFlagRequired("output")

}

//Execute function executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
