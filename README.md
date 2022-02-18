# k8-heimdall

## Description
Derived from the Norse God [Heimdall](https://www.britannica.com/topic/Heimdall) who keeps watch, this is an implementation of event controller that looks out for any event that occurs on the k8 cluster. This is created in an effort to understand and implement the following:

1. Kubernetes client-go package
2. Controller pattern in Kubernetes by use of shared informers and watch

Heimdall watches over the cluster using the kubernetes controller pattern and updates the user in case of any new events on the kubernetes cluster

## Prerequisites

1. Access to a Kubernetes cluster
2. Currently the code assumes that the kubeconfig file will be found in `~/.kube/config`
3. The user should have full access pan cluster for `Event` resource

## How to run

Upon execution the kubernetes events are printed to stdout and the corresponding parsed event fields are logged subsequently

```bash
  $ go run cmd/main.go
```
---
**NOTE**
optionally user can also provide `-kubeconfig` to override default kubeconfig settings
___

## Sample Output

### Server Logs

<img width="1072" alt="Screen Shot 2022-02-18 at 11 15 33 AM" src="https://user-images.githubusercontent.com/3843254/154625347-b4f0ebcb-d2b8-43b3-9924-804e0b743e81.png">

### On Browser (API `/events`)

<img width="787" alt="Screen Shot 2022-02-18 at 11 15 16 AM" src="https://user-images.githubusercontent.com/3843254/154625410-37c5f028-e6b5-4cc8-bcad-6815f060b7c7.png">



