
echo "开始构建"
docker build -t crazyants/my-web:v5 .

echo "构建结束， 开始推送"

docker push crazyants/my-web:v5

echo "推送成功 重启服务"


k delete -f deployment-service.yaml

k apply -f deployment-service.yaml

echo "部署结束"