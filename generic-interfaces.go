package main

import "fmt"

/*
	generic interface -> You can define interfaces that work with generic types.
*/

// Stringer is a generic interface
type Stringer interface {
	String() string
}

type MyType int

func (m MyType) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

func PrintString[T Stringer](t T) {
	fmt.Println(t.String())
}
