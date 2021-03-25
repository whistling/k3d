
### minikube安装

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb

sudo dpkg -i minikube_latest_amd64.deb

minikube version

minikube

minikube ip

minikube tunnel

### minikube启动

minikube start

minikube start --registry-mirror=https://registry.docker-cn.com --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --vm-driver=none

minikube start --vm-driver=virtualbox --docker-env  HTTP_PROXY=http://172.17.20.34:58591 --docker-env  HTTPS_PROXY=http://172.17.20.34:58591 --docker-env NO_PROXY=localhost,127.0.0.1,10.96.0.0/12,172.17.20.34 --registry-mirror=https://rqp9m1n9.mirror.aliyuncs --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --feature-gates=RemoveSelfLink=false

minikube start --vm-driver=virtualbox --registry-mirror=https://rqp9m1n9.mirror.aliyuncs --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --feature-gates=RemoveSelfLink=false

### 面板

minikube dashboard


### minikube插件

minikube addons list

minikube addons enable xxx

minikube addons disable xxx



