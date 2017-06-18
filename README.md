# Setup a new kubernetes cluster on bare metal

Execute the following commands, replace IP-of-master:

```shell
curl https://github.com/alexandervantrijffel/Kubernetes-GPU-Guide/blob/master/scripts/init-master.sh -o init-master.sh
sudo ./init-master.sh <IP-of-master>
mkdir -pv ~/.kube/
cp ~/admin.conf ~/.kube/config
kubectl apply -f https://git.io/weave-kube-1.6
```
Check if everything is working with
```shell
kubectl get pods --all-namespaces
```

Install worker node(s) with these commands,replace Token-of-Master IP-of-master:Port
```shell
curl https://github.com/alexandervantrijffel/Kubernetes-GPU-Guide/blob/master/scripts/init-worker.sh -o init-worker.sh
sudo ./init-worker.sh <Token-of-Master> <IP-of-master>:<Port>
```

## Install a private docker registry

Details at https://github.com/ContainerSolutions/registry-tooling.

Generate certificates and add to kubernetes as secrets
```shell
openssl req -config in.req -newkey rsa:4096 -nodes -sha256 -keyout certs/domain.key -x509 -days 265 -out certs/ca.crt
kubectl create secret generic registry-cert --from-file=./certs/ca.crt 
kubectl create --namespace=kube-system secret generic registry-cert --from-file=./certs/ca.crt 
kubectl create --namespace=kube-system secret generic registry-key --from-file=./certs/domain.key
```
Run https://github.com/alexandervantrijffel/kubernetes-tools/blob/master/dockerregistry/copy-cert.sh on all master and worker nodes (replace /hostfile with /etc/hosts).  

Run https://github.com/ContainerSolutions/registry-tooling/blob/master/reg-tool.sh
```shell
reg-tool.sh install-k8s-reg # (with steps for create-cert and copy-cert commented out) 
```

Install certificate of docker registry with:
```shell
sudo ./reg-tool.sh install-cert --add-host <IP-of-master>
```

Install registry deployment and service
```shell
kubectl apply -f https://raw.githubusercontent.com/alexandervantrijffel/registry-tooling/master/k8s/reg_controller.yaml --record
kubectl apply -f https://github.com/alexandervantrijffel/registry-tooling/blob/master/k8s/reg_service.yaml --record
```

## Install traefik ingress controller for load balancing HTTP requests
Run
```shell
kubectl apply -f https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/traefik-rbac.yaml
kubectl apply -f https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/traefik.yaml
```
Check whether the traefik-ingress-controller is running
```shell
kubectl --namespace=kube-system get pods
```
And can it be accessed on port 80?
```shell
curl <public ip of traefik controller node>
```

List details of the traefik daemonset
```
kubectl describe ds traefik-ingress-controller --namespace=kube-system
```

Web app with ingress routing rules example
```
kubectl apply -f https://raw.githubusercontent.com/containous/traefik/master/examples/k8s/ui.yaml
```
Map the host traefik-ui.minikube to the public ip of the traefik controller mode and the traefik dashboard can be accessed on URL http://traefik-ui.minikube  

Another example app can be found at https://github.com/alexandervantrijffel/kubernetes-tools/tree/master/gowebapp-traefik


# Memory usage of Kubernetes
Bare ubuntu 16.04 server: 197MB in use
Ubuntu 16.04 server with kubernetes master node: 656MB in use (459MB for kubernetes+docker)
Ubuntu 16.04 server with kubernetes worker node: 520MB in use (323MB for kubernetes+docker)
