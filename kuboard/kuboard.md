
### 查看token

kubectl -n kube-system get secret|grep  kuboard-user 

kubectl -n kube-system describe secret kuboard-user-token-wnl5c 
