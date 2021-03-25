
### 文件安装

kubectl apply -f deploy.yaml

测试

kubectl apply -f test-deployment-service.yaml

kubectl apply -f test-ingress.yaml

查看结果
curl http://test-ingress.com/foo

curl http://test-ingress.com/bar



### helm 安装 ingress
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx

helm uninstall ingress-nginx

docker pull pollyduan/ingress-nginx-controller:v0.44.0

k8s.gcr.io/ingress-nginx/controller:v0.44.0

imagePullPolicy: IfNotPresent

测试

kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/usage.yaml






