package config

import (
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

/*
Setup function uses controller-runtime package to extract kubeconfig from the user's system.
This is done by using either KUBECONFIG env var or default to ~/.kube/config

Then using the same config it returns a kubernetes client set.
*/
func Setup() *kubernetes.Clientset {
	config := ctrl.GetConfigOrDie()
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset
}