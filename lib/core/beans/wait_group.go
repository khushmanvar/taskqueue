package beans

import "sync"

var waitGroupInstance *sync.WaitGroup

func WG() *sync.WaitGroup {
	once.Do(func() {
		waitGroupInstance = &sync.WaitGroup{}
	})
	return waitGroupInstance
}