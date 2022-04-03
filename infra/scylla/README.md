https://operator.docs.scylladb.com/stable/generic.html

# Local Dev environment:

minikube start --cpus=6 # start
eval $(minikube docker-env)

# Deploy cert manager:

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.2/cert-manager.yaml

kubectl wait --for condition=established crd/certificates.cert-manager.io crd/issuers.cert-manager.io

kubectl -n cert-manager rollout status deployment.apps/cert-manager-webhook

# Deploy operator:

kubectl apply -f operator.yaml

kubectl wait --for condition=established crd/scyllaclusters.scylla.scylladb.com

kubectl -n scylla-operator rollout status deployment.apps/scylla-operator

Checking logs:
kubectl -n scylla-operator logs deployment.apps/scylla-operator

# Create and initialize a cluster

kubectl create -f cluster.yaml

Verify with:
kubectl -n scylla get ScyllaCluster

Check pods:
kubectl -n scylla get pods

Check cluster:
kubectl -n scylla describe ScyllaCluster simple-cluster

Check each instance logs:
kubectl -n scylla logs simple-cluster-us-east-1-us-east-1a-0 scylla

# Accessing the db

using kubectl:
kubectl exec -n scylla -it simple-cluster-us-east-1-us-east-1a-0 -- cqlsh

inside a pod:

When you create a new Cluster, automatically creates a Service for the clients to use in order to access the Cluster. The service’s name follows the convention <cluster-name>-client.

kubectl -n scylla describe service simple-cluster-client

# Configuring scylla

The operator can take a ConfigMap and apply it to the scylla.yaml configuration file. This is done by adding a ConfigMap to Kubernetes and refering to this in the Rack specification. The ConfigMap is just a file called scylla.yaml that has the properties you want to change in it. The operator will take the default properties for the rest of the configuration.

kubectl create configmap scylla-config -n scylla --from-file=/path/to/scylla.yaml

Wait for it to propagate and restart the cluster:

kubectl rollout restart -n scylla statefulset/simple-cluster-us-east-1-us-east-1a

Configure cassandra-rackdc.properties

kubectl create configmap scylla-config -n scylla --from-file=/tmp/scylla.yaml --from-file=/tmp/cassandra-rackdc.properties -o yaml --dry-run | kubectl replace -f -

# Configure Scylla manager agent

The operator creates a second container for each scylla instance
It serves as a sidecar and it's the main endpoint for interacting with Scylla API. (a controller)

To configure the agent you just create a new secret called scylla-agent-config-secret and populate it with the contents in the scylla-manager-agent.yaml file like this:

kubectl create secret -n scylla generic scylla-agent-config-secret --from-file scylla-manager-agent.yaml

# Scale up

kubectl -n scylla edit ScyllaCluster simple-cluster

To scale up a rack, change the Spec.Members field of the rack to the desired value.

To add a new rack, append the racks list with a new rack. Remember to choose a different rack name for the new rack.

After editing and saving the yaml, check your cluster’s Status and Events for information on what’s happening:

kubectl -n scylla describe ScyllaCluster simple-cluster

# Troubleshooting

check operator:
kubectl -n scylla-operator logs deployment.apps/scylla-operator

check instances:
kubectl -n scylla logs simple-cluster-us-east-1-us-east-1a-0
