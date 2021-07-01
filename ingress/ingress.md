
### 文件安装

kubectl apply -f deploy.yaml


#### 删除权限验证
kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission

测试

kubectl apply -f test-deployment-service.yaml

kubectl apply -f test-ingress.yaml

ip域名绑定
 k get svc -o wide -A |grep ingress-nginx  找到 EXTERNAL-IP

 绑定到host 上
172.23.5.1 test-ingress.com


查看结果
curl http://test-ingress.com/foo

curl http://test-ingress.com/bar



### helm 安装 ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx

helm uninstall ingress-nginx

docker pull pollyduan/ingress-nginx-controller:v0.47.0

docker tag pollyduan/ingress-nginx-controller:v0.47.0 k8s.gcr.io/ingress-nginx/controller:v0.47.0

k8s.gcr.io/ingress-nginx/controller:v0.47.0

imagePullPolicy: IfNotPresent

测试

kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/usage.yaml






