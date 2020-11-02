package main

import (
	"context"
	"log"
	"time"

	"taskqueue/lib/core/beans"
	"taskqueue/proto/gen"

	"google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Failed to connect to server: %v", err)
    }
    defer conn.Close()

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()

    // Example: Enqueue a task
    enqueueTask(ctx, "example-task", "sample payload")

    // Example: Dequeue a task
    // dequeueTask(client, ctx)
}

func enqueueTask(ctx context.Context, name, payload string) {
    res, err := beans.Queue().Enqueue(ctx, &proto.TaskRequest{Name: name, Payload: payload})
    if err != nil {
        log.Fatalf("Failed to enqueue task: %v", err)
    }
    if res.Success {
        log.Println("Task enqueued successfully")
    } else {
        log.Println("Failed to enqueue task")
    }
}

func dequeueTask(client proto.TaskQueueClient, ctx context.Context) {
    task, err := client.Dequeue(ctx, &proto.Empty{})
    if err != nil {
        log.Fatalf("Failed to dequeue task: %v", err)
    }
    log.Printf("Dequeued Task: %v\n", task)
}
