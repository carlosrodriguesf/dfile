package stack

import "sync"

type item[V any] struct {
	prev  *item[V]
	next  *item[V]
	value V
}

type Stack[V any] struct {
	first *item[V]
	last  *item[V]
	size  int
	mutex sync.Mutex
}

func (s *Stack[V]) Size() int {
	return s.size
}

func (s *Stack[V]) Append(value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	next := &item[V]{
		prev:  s.last,
		value: value,
	}

	if s.last != nil {
		s.last.next = next
	}
	if s.first == nil {
		s.first = next
	}
	s.last = next
	s.size++
}

func (s *Stack[V]) AppendList(stack *Stack[V]) {
	stack.mutex.Lock()
	defer func() {
		*stack = Stack[V]{}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if stack.size == 0 {
		return
	}
	if s.size == 0 {
		s.first = stack.first
		s.last = stack.last
		s.size = stack.size
		return
	}

	s.last.next = stack.first
	s.last = stack.last
	s.size += stack.size
}

func (s *Stack[V]) ToArray() []V {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	arr := make([]V, 0, s.size)
	for current := s.first; current != nil; current = current.next {
		arr = append(arr, current.value)
	}
	return arr
}
