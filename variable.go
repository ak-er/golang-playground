package main

import "fmt"

func variable() {
	// Declare variable in go
	var a int // var variableName variableType

	var b int = 23

	// short-hand notation
	c := 44

	fmt.Println(a, b, c)

	const PI = 3.14
	fmt.Println("PI value", PI)

	// exe-1
	var age int
	fmt.Print("Enter Age: ")
	fmt.Scanln(&age)
	fmt.Println("Age is :", age)

	// exe-2
	var height float64
	fmt.Print("Enter height: ")
	fmt.Scanln(&height)
	fmt.Println("Height is :", height)

	// exe-3
	var name string = "akash"
	fmt.Println("Hello,", name)

	// exe-4
	isStudent := false
	fmt.Println(isStudent)

	// exe-5
	const COUNTRY_NAME = "India"
	fmt.Println(COUNTRY_NAME)

	// Write a program that swaps the values of two variables.
	x, y := 10, 20
	x, y = y, x
	fmt.Println(x, y)
}

// func main() {
// 	variable()
// }
