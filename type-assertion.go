package main

import "fmt"

// type assertion -> Typer Assertion is Used to extract the underlying value of an interface.

func TypeAssertion() {
	var i interface{} = 1

	// type assertion
	s, ok := i.(string)
	if !ok {
		fmt.Println("Something went wrong")
	}
	fmt.Println(s)

	// type assertion with check
	if s, ok := i.(string); ok {
		fmt.Println("It's a string", s)
	} else {
		fmt.Println("It's not a string")
	}
}
