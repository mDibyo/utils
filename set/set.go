package set

import "sync"

type exists struct{}

// Set is a thread-safe set implemented on top of a map.
type Set struct{
	m map[interface{}]exists
	l sync.Mutex
}

// Len returns the number of values in the set.
func (s *Set) Len() int {
	return len(s.m)
}

// Add adds a value to the set.
func (s *Set) Add(value interface{}) {
	s.l.Lock()
	defer s.l.Unlock()

	s.m[value] = exists{}
}

// Remove removes a value from the set.
func (s *Set) Remove(value interface{}) {
	s.l.Lock()
	defer s.l.Unlock()

	delete(s.m, value)
}

// Exists returns whether a value is in the set.
func (s *Set) Exists(value interface{}) bool {
	_, ok := s.m[value]
	return ok
}