package main

import "fmt"

func main() {
	const greet = "Hello Golang!"
	// Invalid to reasign  the constant variable
	// greet = "Hello Python"  // get an error => cannot assign to greet (neither addressable nor a map index expression)
	fmt.Println(greet)

	// constant grouping
	const (
		port = 8000
		host = "127.0.0.1"
	)
	fmt.Println(port, host)
}
