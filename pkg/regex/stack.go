package regex

import "errors"

type Stack struct {
	data []byte
}

func (s *Stack) Push(b byte) {
	s.data = append(s.data, b)
}

func (s *Stack) Pop() (byte, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack is empty")
	}
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r, nil
}

func (s *Stack) Top() (byte, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Len() int {
	return len(s.data)
}
