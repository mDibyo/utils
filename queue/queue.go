package queue

import (
	"container/list"
	"sync"
)

// Queue is a thread-safe FIFO queue implemented on top of a linked list.
type Queue struct {
	list  list.List
	mutex sync.Mutex
}

// Len returns the number of values in the queue.
func (q *Queue) Len() int {
	return q.list.Len()
}

// Push inserts a value at the back of the queue.
func (q *Queue) Push(value interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.list.PushBack(value)
}

// Pop removes the value at the front of the queue and returns it.
func (q *Queue) Pop() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.list.Len() == 0 {
		return nil
	}
	return q.list.Remove(q.list.Front())
}

