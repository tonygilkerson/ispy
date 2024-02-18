# ispy

Iot Spy for viewing cluster side info about my IOT stuff

## Local testing

```sh
source .env
go run cmd/ispy/main.go
open http://localhost:8081/
```

## Testing with containers

```sh
# Create VM
KIND_EXPERIMENTAL_PROVIDER=podman
podman machine init --cpus=4 --memory=4000 // adjust resources as needed

# Build container
podman build -t ispy:dev .

# Run
podman run -it --rm -p 8080:8080 localhost/ispy:dev 
open http://localhost:8080/
```

## Testing with kind cluster

This can be use to test the app's chart
This requires port 8080 to be available on the host.

```sh
# Create VM and cluster
podman machine init --cpus=4 --memory=4000 // adjust resources as needed
kind create cluster --config kind-confg.yaml

# Create and load image onto cluster nodes so we don't need a registry
podman build -t ispy:dev . 
podman save -o .temp/ispy.tar localhost/ispy:dev
kind load image-archive .temp/ispy.tar 

# Verify
# If you want to see the image you just loaded on the kind node
$ podman machine ssh
core@localhost:~$  podman exec -it kind-worker /bin/bash
root@kind-worker:/# crictl images  
```

### Install Ingress Controller

>Note installing the ingress controller is optional if you plan to use the service mesh only

```sh
# Install ingress controller
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for controller to be ready
kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=90s

# The following patch will enable allow-snippet-annotations
kubectl -n ingress-nginx patch cm ingress-nginx-controller --type merge --patch-file patch/ingress-nginx-controller-cm-patch.yaml

# The following will patch the nginx ingress to match the KinD `extraPortMappings` used to get traffic into your kind cluster.
kubectl -n ingress-nginx patch svc ingress-nginx-controller --type merge --patch-file patch/ingress-nginx-controller-svc-patch.yaml


# Bounce the pods to pickup the configmap patch
kubectl -n ingress-nginx scale deployment ingress-nginx-controller --replicas=0
kubectl -n ingress-nginx scale deployment ingress-nginx-controller --replicas=1
```

### Deploy and Verify

```sh
# Deploy and verify
helm upgrade -i ispy charts/ispy --set image.repository=localhost/ispy --set image.tag=dev --set ingressClassName=nginx --set domain=127.0.0.1.nip.io 
open https://ispy.127.0.0.1.nip.io:8443/
```
