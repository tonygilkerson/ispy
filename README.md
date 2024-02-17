# ispy

Iot Spy for viewing cluster side info about my IOT stuff

## Local testing

```sh
go run cmd/ispy/main.go
open http://localhost:8080/
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
KIND_EXPERIMENTAL_PROVIDER=podman
podman machine init --cpus=4 --memory=4000 // adjust resources as needed
kind create cluster --config kind-confg.yaml

# Create and load image onto cluster nodes so we don't need a registry
podman save -o .temp/ispy.tar localhost/ispy:dev
kind load image-archive .temp/ispy.tar 

# Verify
# If you want to see the image you just loaded on the kind node
$ podman machine ssh
core@localhost:~$  podman exec -it kind-worker /bin/bash
root@kind-worker:/# crictl images  

# Deploy and verify
helm upgrade -i ispy charts/ispy
open http://ispy.127.0.0.1.nip.io:8080/
```
