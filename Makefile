-include .env
export

CURRENT_DIR=$(shell pwd)

# run service
.PHONY: run
run:
	go run cmd/app/main.go


APP=evrone_api_gateway
CMD_DIR=./cmd

# build for current os
.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# build for linux amd64
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go


# proto
.PHONY: proto-gen
proto-gen:
	./scripts/gen-proto.sh


# go generate
.PHONY: go-gen
go-gen:
	go generate ./...

# generate swagger
.PHONY: swagger-gen
swagger-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs

# run test
.PHONY: test
test:
	go test -v -cover -race ./internal/...

DB_URL := "postgres://postgres:20030505@localhost:5432/farmish?sslmode=disable"

.PHONY: migrate-upy
migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down
# -------------- for deploy --------------
build-image:
	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

push-image:
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

.PHONY: pull-proto-module
pull-proto-module:
	git submodule update --init --recursive

.PHONY: update-proto-module
update-proto-module:
	git submodule update --remote --merge