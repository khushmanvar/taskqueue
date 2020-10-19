package main

import (
	"fmt"
	"sync"
	"taskqueue/lib/core/types"
	"taskqueue/lib/core/worker"
)

func main() {
	queue := types.NewQueue()
	wg := &sync.WaitGroup{}

	// Start workers
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker.Worker(queue, wg)
	}

	// Push tasks
	for i := 0; i < 1000000; i++ {
		queue.Enqueue(fmt.Sprintf("Task %d", i))
	}

	queue.Shutdown()
	wg.Wait()
}
