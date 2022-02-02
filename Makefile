.PHONY: build
build:
	docker-compose build
.PHONY: up
up:
	docker-compose up -d
.PHONY: down
down:
	docker-compose down
.PHONY: exec
exec:
	docker-compose exec goapp sh
.PHONY: log
log:
	docker logs goapp
.PHONY: fmt
fmt:
	docker-compose run --rm goapp go fmt ./...
.PHONY: vet
vet:
	docker-compose run --rm goapp go vet ./...
.PHONY: gotest
gotest:
	docker-compose run --rm goapp go test -v ./...
