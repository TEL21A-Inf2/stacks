package stacks

import "fmt"

func ExampleStackInt_IsEmpty() {
	s1 := NewStackInt()
	fmt.Println(s1.IsEmpty())

	s1.Push(42)
	fmt.Println(s1.IsEmpty())

	s1.Top()
	fmt.Println(s1.IsEmpty())

	s1.Pop()
	fmt.Println(s1.IsEmpty())

	// Output:
	// true
	// false
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

func ExampleStackInt_Top() {
	s1 := NewStackInt()

	v1, e1 := s1.Top()
	fmt.Println(v1, e1)

	s1.Push(42)
	v2, e2 := s1.Top()
	fmt.Println(v2, e2)

	// Output:
	// 0 Top auf leerem Stack aufgerufen
	// 42 <nil>
}

func ExampleStackInt_Push() {
	s1 := NewStackInt()

	s1.Push(42)
	v1, _ := s1.Top()
	fmt.Println(v1)

	s1.Push(25)
	v2, _ := s1.Top()
	fmt.Println(v2)

	s1.Push(38)
	v3, _ := s1.Top()
	fmt.Println(v3)

	// Output:
	// 42
	// 25
	// 38
}
