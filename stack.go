package stacks

type StackInt struct {
	data []int
}

// Konstruktor: Liefert einen neuen Int-Stack.
func NewStackInt() *StackInt {
	return &StackInt{make([]int, 0)}
}
