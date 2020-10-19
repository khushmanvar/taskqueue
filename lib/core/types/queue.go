package types

import "sync"

type Queue struct {
	tasks    chan string
	shutdown chan struct{}
	mu       sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		tasks:    make(chan string, 100),
		shutdown: make(chan struct{}),
	}
}

func (q *Queue) Enqueue(task string) {
	q.tasks <- task
}

func (q *Queue) Dequeue() (string, bool) {
	select {
	case task := <-q.tasks:
		return task, true
	case <-q.shutdown:
		return "", false
	}
}

func (q *Queue) Shutdown() {
	close(q.shutdown)
	close(q.tasks)
}
