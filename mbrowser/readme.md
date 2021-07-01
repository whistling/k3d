
###
1. 生成k8s文件
goctl kube deploy -name redis -namespace adhoc -image redis:6-alpine -o redis.yaml -port 6379

2. 查看
k get pods -n adhoc -o wide

3. sh登录

kubectl run -i --tty --rm cli --image=redis:6-alpine -n adhoc -- sh

4. 连接redis

redis-cli -h redis-svc





