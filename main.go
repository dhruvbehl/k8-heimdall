package main

import (
	"encoding/json"
	"fmt"

	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)

	watchList := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(),
	string("Events"), v1.NamespaceAll, fields.Everything())

	_, controller := cache.NewInformer(watchList, &v1.Event{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) {
			j, _:= json.MarshalIndent(obj, "", "\t")
			fmt.Printf("new event received \n\n[%+v]\n\n", string(j))
		},
	})

	stop := make(chan struct{})
	defer close(stop)

	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}

}