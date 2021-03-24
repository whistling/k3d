
一 minikube安装


curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb

sudo dpkg -i minikube_latest_amd64.deb


minikube version

minikube

二 启动

minikube start

export HTTP_PROXY=http://127.0.0.1:58591; export HTTPS_PROXY=http://127.0.0.1:58591; export ALL_PROXY=socks5://127.0.0.1:51837


export NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,172.17.20.34


minikube start --vm-driver=virtualbox


minikube start --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --vm-driver=none 

minikube start --vm-driver=none --docker-env  HTTP_PROXY=http://127.0.0.1:58591 --docker-env  HTTPS_PROXY=http://127.0.0.1:58591 --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,192.168.99.0/24,172.17.20.34


192.168.99.1
minikube start --vm-driver=virtualbox --docker-env  HTTP_PROXY=http://192.168.99.1:58591 --docker-env  HTTPS_PROXY=http://192.168.99.1:58591 --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12 --registry-mirror=https://rqp9m1n9.mirror.aliyuncs --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --feature-gates=RemoveSelfLink=false

minikube start --vm-driver=virtualbox --docker-env  HTTP_PROXY=http://172.17.20.34:58591 --docker-env  HTTPS_PROXY=http://172.17.20.34:58591 --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,172.17.20.34 --registry-mirror=https://rqp9m1n9.mirror.aliyuncs --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --feature-gates=RemoveSelfLink=false



minikube start --vm-driver=none --docker-env  HTTP_PROXY=http://127.0.0.1:58591 --docker-env  HTTPS_PROXY=http://127.0.0.1:58591 --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12 --registry-mirror=https://rqp9m1n9.mirror.aliyuncs --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --feature-gates=RemoveSelfLink=false


minikube start --vm-driver=none --docker-env  HTTP_PROXY=http://172.17.20.34:58591 --docker-env  HTTPS_PROXY=http://172.17.20.34:58591 --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,172.17.20.34 --feature-gates=RemoveSelfLink=false



minikube start --vm-driver=virtualbox --registry-mirror=https://rqp9m1n9.mirror.aliyuncs --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --feature-gates=RemoveSelfLink=false


minikube start --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --vm-driver=virtualbox 



./minikube start --vm-driver=virtualbox  --image-mirror-country=cn --iso-url=/home/hanzq/App/minikube-v1.17.0.iso --registry-mirror=https://rqp9m1n9.mirror.aliyuncs.com



三 部署应用

minikube dashboard

例如下面的命令 

kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.10

替换后为 

kubectl create deployment hello-minikube --image=registry.aliyuncs.com/google_containers/echoserver:1.10 即可。

kubectl describe pod nginx-deployment-7848d4b86f-4q7vv




四 穷人版balance

172.17.5.0/28

kubectl get pods --all-namespaces

minikube tunnel

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.5/manifests/namespace.yaml

kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)"

kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey='(openssl rand -base64 128)'

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.5/manifests/metallb.yaml


kubectl apply -f metallb-configMap.yaml


kubectl get pods -n metallb-system



kubectl create deployment nginx --image nginx



kubectl expose deployment/nginx --type="LoadBalancer" --port 8080



kubectl logs controller-57f648cb96-fm645 -n metallb-system



五 插件

minikube addons list



六 运行外部访问

kubectl proxy  --port=8001 --address='10.0.0.2' --accept-hosts='^.*' &



 七 kind

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.10.0/kind-linux-amd64
chmod +x ./kind
mv ./kind /usr/bin


kubectl apply -f recommended_own.yaml


http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login


kubectl apply -f admin-user-role-binding.yaml


kubectl -n kube-system describe secret '(kubectl -n kube-system get secret | grep admin | awk '{print $1}')' |grep admin-user-token

kubectl -n kube-system get secret|grep admin

kubectl -n kube-system describe secret admin-token-6pvrv 


