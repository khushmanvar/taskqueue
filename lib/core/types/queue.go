package types

import "sync"

type Queue struct {
	Tasks    chan string
	shutdown chan struct{}
	mu       sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		Tasks:    make(chan string, 100),
		shutdown: make(chan struct{}),
	}
}

func (q *Queue) Enqueue(task string) {
	q.Tasks <- task
}

func (q *Queue) Dequeue() (string, bool) {
	select {
	case task := <-q.Tasks:
		return task, true
	case <-q.shutdown:
		return "", false
	}
}

func (q *Queue) Shutdown() {
	close(q.shutdown)
	close(q.Tasks)
}
