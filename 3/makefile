server:
	go run cmd/server/main.go

proto:
	protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative \
    --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
    internal/proto/*.proto

.PHONY: server proto