package Queue

import "sync"

type Queue struct {
	queue []interface{}
	len   int
	lock  *sync.Mutex
}

func New() *Queue {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0
	queue.lock = new(sync.Mutex)
	return queue
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}
