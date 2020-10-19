package worker

import (
	"log"
	"sync"
	"taskqueue/lib/core/types"
	"time"
)

func Worker(queue *types.Queue, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := queue.Dequeue()
		if !ok {
			return
		}
		processTask(task)
	}
}

func processTask(task string) {
	log.Printf("Processing task: ", task)
	time.Sleep(1 * time.Second)
}
