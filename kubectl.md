
###kubectl安装
curl -LO "https://dl.k8s.io/release/'(curl -L -s https://dl.k8s.io/release/stable.txt)'/bin/linux/amd64/kubectl"

sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl


### 启动服务
kubectl create deployment nginx --image nginx

kubectl expose deployment/nginx --type="LoadBalancer" --port 80

### 拉取镜像
kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.10
替换后为
kubectl create deployment hello-minikube --image=registry.aliyuncs.com/google_containers/echoserver:1.10

docker tag docker.io/kubernetesui/dashboard:v2.2.0 k8s.gcr.io/kubernetesui/dashboard:v2.2.0

### 穷人版balance

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/main/manifests/namespace.yaml

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/main/manifests/metallb.yaml

kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey='(openssl rand -base64 128)'

kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"

kubectl get pods -n metallb-system

使用
kubectl apply -f configmap.yaml

配置的ip段和 node的ip端前两位一致
kubectl get nodes -o wide

kubectl logs controller-57f648cb96-fm645 -n metallb-system

https://github.com/metallb/metallb/tree/main/manifests


### 常用命令
kubectl get nodes

kubectl get pods -A
kubectl get pods --all-namespaces

kubectl get deployments

kubectl get svc


kubectl describe pod kubernetes-dashboard-9f9799597-fmlhq  -n kubernetes-dashboard

kubectl logs controller-fb659dc8-lkqjx -n metallb-system

kubectl apply -f dashboard_own.yaml

kubectl delete -f recommended.yaml

kubectl delete pods dashboard-metrics-scraper-79c5968bdc-v7g2c -n kubernetes-dashboard

kubectl port-forward service/nginx-deployment 8080:80


kubectl get svc nginx-deployment -o yaml

kubectl get nodes -o yaml | grep ExternalIP -C 1

### 代理

export HTTP_PROXY=http://127.0.0.1:58591; export HTTPS_PROXY=http://127.0.0.1:58591; export ALL_PROXY=socks5://127.0.0.1:51837


export HTTP_PROXY=http://172.17.224.1:7890; export HTTPS_PROXY=http://172.17.224.1:7890; export ALL_PROXY=socks5://127.0.0.1:1080

export NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,172.17.20.34



