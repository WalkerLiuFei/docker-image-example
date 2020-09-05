#!/bin/bash
echo "Start build image"
docker build --tag walkerliu/http-server-example:latest .
echo "Build image done"
docker image ls | grep http-server
echo "push to remote docker hub...."
docker push walkerliu/http-server-example:latest
echo "run "

echo  "--------------run dokcer image in kubernetes，https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/--------------------"


echo "kubectl start the pod runing http-server in docker and kubernetes cluster"
kubectl apply -f ./deploy.yaml

echo "expose a port through with a service 'http-server-service'"
kubectl expose deployment http-server-deployment --type=LoadBalancer --name=http-server-service

# 这个时候才会将 端口暴露出来，但是不是通过 docker那样直接expose的端口，具体的端口需要从下面命令看到，
# 具体的端口是在NodePort 那个属性值

echo "describe depoloyment"
kubectl describe service http-server-service

## 在 mac 和 windows