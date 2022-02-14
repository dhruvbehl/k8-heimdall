# k8-heimdall

## Description
Derived from the Norse God [Heimdall](https://www.britannica.com/topic/Heimdall) who keeps watch, this is an implementation of event controller that looks out for any event that occurs on the k8 cluster 

Heimdall watches over the cluster using the kubernetes controller pattern and updates the user in case of any new events on the kubernetes cluster

## Prerequisites

1. Access to a Kubernetes cluster
2. Currently the code assumes that the kubeconfig file will be found in `~/.kube/config`
3. The user should have full access pan cluster for `Event` resource

## How to run

```bash
  $ go run main.go
```
