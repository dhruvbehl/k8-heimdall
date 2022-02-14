package main

import (
	"flag"
	"os"

	"github.com/dhruvbehl/k8-heimdall/config"
	"github.com/dhruvbehl/k8-heimdall/pkg/events"
)

func main() {
	kubeconfig := flag.String("kconfig", "", "custom kubeconfig path to be given by user [optional], by default the program uses the default kubeconfig in the env var or ~/.kube/config")
	flag.Parse()
	if *kubeconfig != "" {
		os.Setenv("KUBECONFIG", *kubeconfig)
	}
	clientset := config.Setup()
	events.WatchEvents(clientset)
}