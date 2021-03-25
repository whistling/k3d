### helm 安装

curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash


helm repo add stable  https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts/


helm repo add nginx-stable https://helm.nginx.com/stable

docker pull pollyduan/ingress-nginx-controller:v0.44.0

helm repo update

helm search repo ingress

helm install pkslow-ingress azure/nginx-ingress

k8s.gcr.io/ingress-nginx/controller:v0.44.0


kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.44.0/deploy/static/provider/baremetal/deploy.yaml