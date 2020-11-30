package beans

import (
	"log"

	"google.golang.org/grpc"
)

var (
	connInstance *grpc.ClientConn
)

func Conn() *grpc.ClientConn {
	once.Do(func() {
		connInstance, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer connInstance.Close()
	})
	return connInstance
}