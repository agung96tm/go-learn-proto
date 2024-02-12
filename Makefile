dependencies: ## generate dependencies
	go mod download

proto_generate: ## generate proto file
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative resources/protos/users.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative resources/protos/posts.proto

runapi: ## run application as api (client)
	go run main.go runapi

runrpc: ## run application as rpc (server)
	go run main.go runrpc

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help