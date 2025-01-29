package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

/*
func addingTwo(fn func(a int) int) int {
	return fn(4)
}
*/

func addingTwo() func(a int) int {
	return func(a int) int {
		return a + 2
	}
}

func main() {
	a := 10
	b := 34
	result := addition(a, b)
	fmt.Println(result)

	/*
		fn := func(a int) int {
			return a + 2
		}

		x := addingTwo(fn)
		fmt.Println(x)
	*/

	x := addingTwo()(5)
	fmt.Println(x)
}
