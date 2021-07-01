### 安装cilium

curl -L --remote-name-all https://github.com/cilium/cilium-cli/releases/latest/download/cilium-linux-amd64.tar.gz{,.sha256sum}
sha256sum --check cilium-linux-amd64.tar.gz.sha256sum


sudo tar xzvfC cilium-linux-amd64.tar.gz /usr/local/bin

cilium install

cilium

cilium service list

cilium hubble enable

cilium hubble


### 检查

kubectl apply -f https://raw.githubusercontent.com/cilium/cilium/1.7.2/examples/kubernetes/connectivity-check/connectivity-check.yaml


### hubble

helm repo add cilium https://helm.cilium.io/


helm install cilium cilium/cilium --version 1.7.3 \
--namespace kube-system \
--set global.cni.chainingMode=aws-cni \
--set global.masquerade=false \
--set global.tunnel=disabled \
--set global.nodeinit.enabled=true




cilium hubble port-forward&

hubble status

hubble observe



cilium hubble enable --ui


cilium hubble ui

