package stacks

import "fmt"

func ExampleStackInt_IsEmpty() {
	s1 := NewStackInt()
	fmt.Println(s1.IsEmpty())

	// Output:
	// true
}
