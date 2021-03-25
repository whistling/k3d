
### 安装

curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.10.0/kind-linux-amd64

chmod +x ./kind

mv ./kind /usr/bin

kind create cluster



### 启动面板

kubectl apply -f dashboard.yaml

kubectl proxy

kubectl proxy  --port=8001 --address='10.0.0.2' --accept-hosts='^.*' &

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login


### 授权&获取token

kubectl apply -f admin-user-role-binding.yaml

kubectl -n kube-system get secret|grep admin

kubectl -n kube-system describe secret admin-token-6pvrv 




kubectl -n kube-system describe secret '(kubectl -n kube-system get secret | grep admin | awk '{print $1}')' |grep admin-user-token

