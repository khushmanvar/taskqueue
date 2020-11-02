.PHONY: proto server worker

proto:
	sh scripts/generate_proto.sh

server:
	go run cmd/server/main.go

worker:
	go run cmd/worker/main.go