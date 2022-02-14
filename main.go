package main

import (
	"flag"
	"os"

	"github.com/dhruvbehl/k8-heimdall/config/kubeconfig"
	"github.com/dhruvbehl/k8-heimdall/events/watch"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "custom kubeconfig path to be given by user [optional], by default the program uses the default kubeconfig in the env var or ~/.kube/config")
	flag.Parse()
	if *kubeconfig != "" {
		os.Setenv("KUBECONFIG", *kubeconfig)
	}
	clientset := kubeconfig.Setup()
	watch.watchEvents(clientset)
}