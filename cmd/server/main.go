package main

import (
	"log"
	"net"

	proto "taskqueue/proto/gen"

	"google.golang.org/grpc"
)

type qserver struct {
    proto.UnimplementedTaskQueueServer
}

func generateTaskID() string {
    return "task-12345"
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterTaskQueueServer(grpcServer, &qserver{})

    log.Println("Server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
