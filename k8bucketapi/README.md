# Using Kubernetes and Minikube

For Local Use:

--------------

## The current steps below applies

1.start minikube

```shell
minikube start
```

2.Enable the ingress add-on for Minikube

```shell
minikube addons enable ingress
```

3.Create a ServiceAccount and associate it with the ClusterRole, use a ClusterRoleBinding

```shell
kubectl create serviceaccount -n kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
```

4.Initialize Helm

```shell
helm init --history-max 200 --service-account tiller
```

5.Confirm tiller works

```shell
kubectl --namespace kube-system get pods | grep tiller
```

6.Bind the prometheus service account

```shell
kubectl apply -f crbinding.yaml
```

7.Validate chart template ###

```shell
helm template k8bucketapi
```

OR

```shell
helm lint k8bucketapi (Chart)
```

8.Do a dry run

```shell
helm install --name api-release --dry-run --debug k8bucketapi
```

9.Install helm chart

```shell
helm install --name api-release k8bucketapi
```

10.Visit a service

```shell
minikube service api-release-k8bucketapi --url
minikube service api-release-prometheus-server --url
minikube service api-release-grafana --url
```

11.Describe the ingress (k8bucketapi, prometheus-server, grafana)

```shell
kubectl describe ing api-release-k8bucketapi
kubectl describe ing api-release-prometheus-server
kubectl describe ing api-release-grafana
```

12.Update your /etc/hosts file to route requests

```shell
echo "$(minikube ip) gobucketapi.local monitoring.local" | sudo tee -a /etc/hosts
```

13.Visit the urls to see the services running

```text
http://api.local
http://prometheus.local
http://grafana.local
```

## Upgrade chart

```shell
helm upgrade api-release k8bucketapi
```

## Delete api-release

```shell
helm del --purge api-release
```

For Production Use

-------------------;

### All the steps above applies with the exception of the nginx-ingress-controller

Replace the steps above with this

2.create the “mandatory” resources for Nginx Ingress in your cluster.

```shell
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
```

2a. create the ingress-nginx ingress controller service

```shell
kubectl apply -f cloud-generic.yaml
```
