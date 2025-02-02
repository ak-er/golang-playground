package main

import "fmt"

func printSlice[T any](slice []T) {
	for _, v := range slice {
		fmt.Println(v)
	}
}

func main() {
	/*
		// Basic generics
		slice := []int{1, 2, 3, 4, 5}
		printSlice(slice)
		stringSlice := []string{"python", "golang", "java"}
		printSlice(stringSlice)
	*/

	/*
		// constraints.Order
		fmt.Println(Max(3.342, 2.234234))
		fmt.Println(Max(34, 56))
		fmt.Println(Max("b", "a"))
		fmt.Println(Max("akash", "vikas"))
	*/

	/*
		// generics types

		stack := Stack[int]{
			elements: []int{1, 2},
		}
		fmt.Println(stack) // 1, 2
		stack.Push(3)
		stack.Push(4)
		fmt.Println(stack) // 1, 2, 3, 4
		stack.Pop()
		fmt.Println(stack) // 1, 2, 3

		stringStack := Stack[string]{
			elements: []string{"golang", "python"},
		}
		fmt.Println(stringStack) // golang, python
		stringStack.Push("java")
		stringStack.Push("c++")
		fmt.Println(stringStack) // golang, python, java, c++
		stringStack.Pop()
		fmt.Println(stringStack) // golang, python, java
	*/

	/*
		// Custom constraints
		slice := []int{1, 2, 3, 4, 5}
		fmt.Println(Sum(slice))
		floatSlice := []float64{2.21, 1.23, 12.232}
		fmt.Println(Sum(floatSlice))
	*/

	/*
		// Type inference
		fmt.Println(Identity(34))
		fmt.Println(Identity("golang"))
	*/

	// generics interfaces
	var m MyType = 23
	PrintString(m)
}
