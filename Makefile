proto:
	protoc -I . pkg/protos/*.proto  --proto_path=pkg/protos --go_out=. --go-grpc_out=.
