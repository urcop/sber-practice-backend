BUILD_DIR = build

run:
	go run -v ./cmd/app/main.go serve

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build  -o ${BUILD_DIR}/app ./cmd/app

swagger:
	swag init --parseDependency -g cmd/app/main.go --output=./api

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/grpc/schema/*.proto

install-tools:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/golang/mock/mockgen@v1.6.0

swagger:
	swag init --parseDependency -g cmd/app/main.go --output=./api

run-dev:
	docker compose up -d --build