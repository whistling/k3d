
# Example:
#   make run

#   make build
#   make build
#   make restart
.PHONY: run build push restart

V ?= v4

test:
	echo $(V)

run:
	$(info tag version: $(V))
	make build
	make push
	make restart

build:
	docker build -t crazyants/my-web:$(V) .

push:
	docker push crazyants/my-web:$(V)

restart:
	kubectl apply -f deployment-service.yaml
