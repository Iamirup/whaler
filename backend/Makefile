protoc-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

gen-blog-proto:
	protoc \
    --proto_path=protobuf "protobuf/blog.proto" \
	--go_out=common/genproto/blog --go_opt=paths=source_relative \
  	--go-grpc_out=common/genproto/blog --go-grpc_opt=paths=source_relative