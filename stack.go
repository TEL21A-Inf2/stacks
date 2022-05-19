package stacks

import "errors"

type StackInt struct {
	data []int
}

// Konstruktor: Liefert einen neuen Int-Stack.
func NewStackInt() *StackInt {
	return &StackInt{make([]int, 0)}
}

// Liefert true, falls der Stack leer ist.
func (s StackInt) IsEmpty() bool {
	return len(s.data) == 0
}

// Fügt ein Element zum Stack hinzu.
func (s *StackInt) Push(value int) {
	s.data = append(s.data, value)
}

// Entfernt das oberste Element vom Stack und liefert es.
func (s *StackInt) Pop() (int, error) {
	top, error := s.Top()
	if error != nil {
		return 0, errors.New("Pop auf leerem Stack aufgerufen")
	}
	s.data = s.data[:len(s.data)-1]
	return top, nil
}

// Liefert das oberste Element vom Stack.
func (s StackInt) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Top auf leerem Stack aufgerufen")
	}
	return s.data[len(s.data)-1], nil
}
