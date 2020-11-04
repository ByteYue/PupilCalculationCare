package stack

import (
	"sync"
)

type Stack struct {
	elements []interface{}
	lock     sync.RWMutex
}

func (s *Stack) New() *Stack {
	s.elements = []interface{}{}
	return s
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	index := len(s.elements) - 1
	item := s.elements[index]
	s.elements = s.elements[0:index]
	return item
}

func (s *Stack) Top() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	index := len(s.elements) - 1
	return s.elements[index]
}

func (s *Stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.elements)
}