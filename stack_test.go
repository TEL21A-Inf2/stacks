package stacks

import "fmt"

func ExampleStackInt_IsEmpty() {
	s1 := NewStackInt()
	fmt.Println(s1.IsEmpty())

	s1.Push(42)
	fmt.Println(s1.IsEmpty())

	s1.Pop()
	fmt.Println(s1.IsEmpty())

	// Output:
	// true
	// false
	// true
}

func ExampleStackInt_Pop() {
	s1 := NewStackInt()

	v1, e1 := s1.Pop()
	fmt.Println(v1, e1)

	s1.Push(42)
	v2, e2 := s1.Pop()
	fmt.Println(v2, e2)

	// Output:
	// 0 Pop auf leerem Stack aufgerufen
	// 42 <nil>
}
