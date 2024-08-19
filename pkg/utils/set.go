package utils

import (
	"sync"
)

// SetElement the type of the Set
type SetElement interface{}

// Set the set of Items
type Set struct {
	items map[SetElement]bool
	lock  sync.RWMutex
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *Set) Add(t SetElement) *Set {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[SetElement]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *Set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = make(map[SetElement]bool)
}

// Delete removes the SetElement from the Set and returns Has(SetElement)
func (s *Set) Delete(item SetElement) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the SetElement
func (s *Set) Has(item SetElement) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.items[item]
	return ok
}

// Items returns the SetElement(s) stored
func (s *Set) Items() []SetElement {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := []SetElement{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
