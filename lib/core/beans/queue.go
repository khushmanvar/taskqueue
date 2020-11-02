package beans

import (
	"log"
	"sync"
	proto "taskqueue/proto/gen"

	"google.golang.org/grpc"
)

var (
	instance proto.TaskQueueClient
	once sync.Once
)

func Queue() proto.TaskQueueClient {
	once.Do(func() {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		instance = proto.NewTaskQueueClient(conn)
	})
	return instance
}