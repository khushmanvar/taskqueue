package beans

import (
	"sync"
	proto "taskqueue/proto/gen"
)

var (
	queueInstance proto.TaskQueueClient
	once sync.Once
)

func Queue() proto.TaskQueueClient {
	once.Do(func() {
		queueInstance = proto.NewTaskQueueClient(Conn())
	})
	return queueInstance
}