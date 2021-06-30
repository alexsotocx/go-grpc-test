docker_build_server:
	docker build -f dockerfile-server -t minddocasoto/grpc_test_server .

docker_build_client:
	docker build -f dockerfile-client -t minddocasoto/grpc_test_client .

run_server:
	SERVER_NAME='SERVER 1' go run server/server.go

run_client:
	SERVER_ADDRESS="localhost" \
	SERVER_PORT="8000" \
	go run client/client.go

compile:
	protoc proto/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
