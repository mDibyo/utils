package queue

import (
	"container/list"
	"sync"
)

// Queue is a thread-safe FIFO queue implemented on top of a linked list.
// The zero-value is a valid empty queue.
type Queue struct {
	l list.List
	m sync.Mutex
}

// Len returns the number of values in the queue.
func (q *Queue) Len() int {
	return q.l.Len()
}

// Push inserts a value at the back of the queue.
func (q *Queue) Push(value interface{}) {
	q.m.Lock()
	defer q.m.Unlock()

	q.l.PushBack(value)
}

// Pop removes the value at the front of the queue and returns it.
func (q *Queue) Pop() interface{} {
	q.m.Lock()
	defer q.m.Unlock()

	if q.l.Len() == 0 {
		return nil
	}
	return q.l.Remove(q.l.Front())
}

