### 安装

curl -L https://git.io/getLatestIstio | sh -

sudo mv istio-1.9.2/ /usr/local/

sudo ln -s /usr/local/istio-1.9.2/bin/istioctl  istioctl

istioctl


### 安装istio

istioctl install --set profile=demo -y








