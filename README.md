# flashy

minikube start --cpus=6
eval $(minikube docker-env)

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.1/cert-manager.yaml
kubectl wait --for condition=established crd/certificates.cert-manager.io crd/issuers.cert-manager.io
kubectl -n cert-manager rollout status deployment.apps/cert-manager-webhook

kubectl apply -f operator.yaml
kubectl wait --for condition=established crd/scyllaclusters.scylla.scylladb.com
kubectl -n scylla-operator rollout status deployment.apps/scylla-operator

kubectl create -f cluster.yaml

kubectl -n scylla get ScyllaCluster

kubectl -n scylla get pods

kubectl -n scylla describe ScyllaCluster simple-cluster

kubectl -n scylla logs simple-cluster-us-east-1-us-east-1a-0 scylla

kubectl create configmap scylla-config -n scylla --from-file=scylla.yaml -o yaml --dry-run=client | kubectl replace -f -
