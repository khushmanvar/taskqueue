.PHONY: proto server worker benchmark

proto:
	sh scripts/generate_proto.sh

server:
	go run cmd/server/main.go

worker:
	go run cmd/worker/main.go

benchmark:
	go test -bench=./benchmarks