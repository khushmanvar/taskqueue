package main

import (
	"fmt"
	"taskqueue/lib/core/beans"
	"taskqueue/lib/core/types"
	"testing"
)

var (
	wN = 3
	tN = 10
)

func BenchmarkWorkerAddTask(b *testing.B) {
    queue := types.NewQueue()
    wg := beans.WG()

    // Start workers
	for i := 0; i < wN; i++ {
		wg.Add(1)
		go types.Worker(queue, wg)
	}

    b.ResetTimer()

    // Push tasks
	for i := 0; i < tN; i++ {
		queue.Enqueue(fmt.Sprintf("Task %d", i))
	}

    queue.Shutdown()
	wg.Wait()
}