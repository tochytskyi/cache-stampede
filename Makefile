.PHONY: build
build:
	docker-compose build --no-cache --force-rm

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: prune
prune:
	make down && docker rmi treatfield-api

.PHONY: restart
restart:
	docker-compose restart

.PHONY: rebuild
rebuild:
	docker exec treatfield-api go build -o /main /go/src/treatfield-api && docker restart treatfield-api

.PHONY: ps
ps:
	docker ps -a | grep treatfield-api

.PHONY: psi
psi:
	docker images | grep treatfield-api

.PHONY: bash
bash:
	docker exec -it treatfield-api sh

.PHONY: test
test:
	docker exec treatfield-api go test ./...
