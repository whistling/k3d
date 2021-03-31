
### k8s 安装

helm repo add dapr https://dapr.github.io/helm-charts/
helm repo update
# See which chart versions are available
helm search repo dapr --devel --versions

helm upgrade --install dapr dapr/dapr \
--version=1.0.1 \
--namespace dapr-system \
--create-namespace \
--wait


kubectl get pods --namespace dapr-system


###高可用安装
helm upgrade --install dapr dapr/dapr \
--version=1.0.1 \
--namespace dapr-system \
--create-namespace \
--set global.ha.enabled=true \
--wait

### 卸载

helm uninstall dapr --namespace dapr-system


### 本机安装

wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash











