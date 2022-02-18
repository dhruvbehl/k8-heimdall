package main

import (
	"flag"
	"os"

	"github.com/dhruvbehl/k8-heimdall/config"
	"github.com/dhruvbehl/k8-heimdall/pkg/events"
	"github.com/rs/zerolog/log"
)

func main() {
	kubeconfig := flag.String("kconfig", "", "custom kubeconfig path to be given by user [optional], by default the program uses the default kubeconfig in the env var or ~/.kube/config")
	httpAddress := flag.String("address", ":8080", "http address of the server")
	flag.Parse()

	go func() {
		engine := events.Setup()
		if err := engine.Run(*httpAddress); err!=nil {
			log.Fatal().Err(err).Msgf("unable to initiate server at [%v]", *httpAddress)
		}
	}()

	if *kubeconfig != "" {
		os.Setenv("KUBECONFIG", *kubeconfig)
	}

	// TODO: Add gin rest server as a go routine so it doesn't block
	clientset := config.Setup()
	events.WatchEvents(clientset)
}