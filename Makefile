build:
	docker build -t grpc-play .

.PHONY: run

run:
	SERVER_NAME='SERVER 1' go run server/server.go

compile:
	protoc proto/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
