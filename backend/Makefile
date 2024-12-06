protoc-install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

gen-most_likes_articles:
	@protoc \
    --proto_path=protobuf "protobuf/most_likes_articles.proto" \
	--go_out=common/genproto/most_likes_articles --go_opt=paths=source_relative \
  	--go-grpc_out=common/genproto/most_likes_articles --go-grpc_opt=paths=source_relative

gen-top_authors:
	@protoc \
    --proto_path=protobuf "protobuf/top_authors.proto" \
	--go_out=common/genproto/top_authors --go_opt=paths=source_relative \
  	--go-grpc_out=common/genproto/top_authors --go-grpc_opt=paths=source_relative
