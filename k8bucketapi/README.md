# Using Kubernetes and Minikube

1. start minikube
minikube start

2. Enable the ingress add-on for Minikube.
minikube addons enable ingress

3. Create a ServiceAccount and associate it with the ClusterRole, use a ClusterRoleBinding
kubectl create serviceaccount -n kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller

4. Initialize Helm
helm init --history-max 200 --service-account tiller

5. Confirm tiller works
kubectl --namespace kube-system get pods | grep tiller

6. Bind the prometheus service account
kubectl apply -f crbinding.yaml

7. Validate chart template ###
helm template k8bucketapi || helm lint k8bucketapi (Chart)

8. Do a dry run
helm install --name api-release --dry-run --debug k8bucketapi

9. Install helm chart
helm install --name api-release k8bucketapi

10. Visit a service
minikube service api-release-k8bucketapi --url
minikube service api-release-prometheus-server --url
minikube service api-release-grafana --url

11. Describe the ingress (k8bucketapi, prometheus-server, grafana)
kubectl describe ing api-release-k8bucketapi
kubectl describe ing api-release-prometheus-server
kubectl describe ing api-release-grafana

12. Update our /etc/hosts file to route requests
echo "$(minikube ip) gobucketapi.local monitoring.local" | sudo tee -a /etc/hosts

## upgrade chart

helm upgrade api-release k8bucketapi

## delete api-release

helm del --purge api-release 
