docker tag gowebapp:latest kube-registry.kube-system.svc.cluster.local:31000/gowebapp:latest
docker push kube-registry.kube-system.svc.cluster.local:31000/gowebapp:latest

