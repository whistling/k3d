
# Example:
#   make run

#   make build
#   make build
#   make restart
.PHONY: build

run:
	make build
	make push
	make restart

build:
	docker build -t crazyants/my-web:v4 .

push:
	docker push crazyants/my-web:v4

restart:
	kubectl apply -f deployment-service.yaml
