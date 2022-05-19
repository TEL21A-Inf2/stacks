package stacks

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
