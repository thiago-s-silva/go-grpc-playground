.PHONY: default run build test docs clean

# Tasks
default: run

run:
	@go run cmd/grpcServer/main.go
build:
	@go build -o $(APP_NAME) main.go
protoc:
	@protoc --go_out=. --go-grpc_out=. proto/course_category.proto
evans:
	@evans --proto ./proto/course_category.proto repl