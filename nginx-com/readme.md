
查看ip
minikube ip

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb



1. minikube安装

curl -Lo minikube https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/releases/v1.17.1/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/



2. kubectl安装


curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"


sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl




3. 启动

minikube start --vm-driver=virtualbox --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --registry-mirror=https://rqp9m1n9.mirror.aliyuncs.com  --feature-gates=RemoveSelfLink=false




minikube start --vm-driver=virtualbox --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --iso-url=https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v1.7.3.iso  --registry-mirror=https://hub-mirror.c.163.com



minikube start --vm-driver=virtualbox --image-mirror-contry=cn --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --iso-url=https://kubernetes.oss-cn-hangzhou.aliyuncs.com/minikube/iso/minikube-v1.7.3.iso  --registry-mirror=https://hub-mirror.c.163.com




4. virtualbox 

sudo apt install --reinstall linux-headers-5.4.70-amd64-desktop virtualbox-dkms dkms





5. k8s 部署nginx 入门

kubectl create deployment my-nginx --image=nginx
kubectl expose deployment my-nginx --type=LoadBalancer --port=8080



6. 查看错误信息
kubectl describe pod fail-1036623984-hxoas







kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.3/manifests/namespace.yaml

kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.9.3/manifests/metallb.yaml


kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey='(openssl rand -base64 128)'



kubectl delete -f https://raw.githubusercontent.com/metallb/metallb/v0.9.3/manifests/namespace.yaml

kubectl delete -f https://raw.githubusercontent.com/metallb/metallb/v0.9.3/manifests/metallb.yaml




kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey=1yhVWYQnwp1O+mCZqQ0Atbb2+eNT73pPpE+Dg8/53WM3etEAlS0KuFF5IrIjK2NcS5L97oj+G39z9hRQRH+R0h+CfV0qptY86kktfpSj9Vnw4TZXkeV53iibtLM252wcZZWuLkbu44f2MDASko3GMNcsRx5cVmrevnOQUs2MM78=




kubectl logs controller-fb659dc8-lkqjx -n metallb-system





lugin=cni --bootstrapper=localkube --feature-gates=CustomResourceValidation=true

