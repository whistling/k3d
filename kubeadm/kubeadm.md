


sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl

sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg



echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list



sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl












### 错误修复

1.  detected "cgroupfs" as the Docker cgroup driver.

sudo mkdir /etc/docker
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

2.    [ERROR Swap]: running with swap on is not supported. Please disable swap

```
sudo swapoff -a

一、不重启电脑，禁用启用swap，立刻生效

# 禁用命令

sudo swapoff -a
# 启用命令

sudo swapon -a
 
# 查看交换分区的状态
sudo free -m


二、重新启动电脑，永久禁用Swap
# 把根目录文件系统设为可读写
sudo mount -n -o remount,rw /
# 用vi修改/etc/fstab文件，在swap分区这行前加 # 禁用掉，保存退出

vi /etc/fstab

 
重新启动电脑，使用free -m查看分区状态
reboot
sudo free -m

```

