package stack

import "errors"

type Stack interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
}

type Storage struct {
	stackStore []int
	pointer    int
}

func (s *Storage) Push(value int) {
	s.stackStore = append(s.stackStore, value)
	s.pointer = len(s.stackStore) - 1
}

func (s *Storage) Pop() (int, error) {
	if s.pointer == -1 {
		return 0, errors.New("stack Empty")
	}
	value := s.stackStore[s.pointer]
	s.stackStore = s.stackStore[:s.pointer]
	s.pointer = len(s.stackStore) - 1
	return value, nil
}

func (s *Storage) Peek() (int, error) {
	if s.pointer == -1 {
		return 0, errors.New("stack Empty")
	}

	return s.stackStore[s.pointer], nil
}

func NewStack() *Storage {
	return &Storage{
		stackStore: make([]int, 0),
		pointer:    -1,
	}
}
