all: proto mongo
	@go run ./cmd

proto:
	@protoc --go-grpc_out=require_unimplemented_servers=false:. --go_out=. proto/notification.proto

mongo:
	@docker run --name mongodb -d -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=qwerty123 -p 27017:27017 mongo

.PHONY: proto