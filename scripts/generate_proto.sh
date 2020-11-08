protoc --proto_path=proto/source \
  --plugin=protoc-gen-go=/home/khush/go/bin/protoc-gen-go \
  --plugin=protoc-gen-go-grpc=/home/khush/go/bin/protoc-gen-go-grpc \
  --go_out=paths=source_relative:proto/gen \
  --go-grpc_out=paths=source_relative:proto/gen \
  proto/source/*.proto