package main

import (
	"math"
)

func divide(a int, b int) (int, int) {
	quotient := a / b
	reminder := a % b
	return quotient, reminder
}

// func main() {
// 	q, r := divide(34, 12)
// 	fmt.Println(q, r)
// }

// anonymous and closure function

// func main() {
// 	// anonymous function
// 	square := func(x int) int {
// 		return x * x
// 	}
// 	fmt.Printf("square of %d is %d\n", 4, square(4))

// 	// closure function example - counter
// 	increment := func() func() int {
// 		count := 0
// 		return func() int {
// 			count++
// 			return count
// 		}
// 	}

// 	inc := increment()
// 	fmt.Println(inc()) // 1
// 	fmt.Println(inc()) // 2
// 	fmt.Println(inc()) // 3
// 	fmt.Println(inc()) // 4
// 	fmt.Println(inc()) // 5
// }

/*

## Exercises:
--Create a function greetUser that takes a name as a parameter and prints "Hello, [Name]!".
--Write a function multiply that takes two integers as input and returns their product.
--Implement a function isEven that takes an integer as input and returns true if it’s even, false otherwise.
--Create a function calculate that takes two integers and an operation (+, -, *, /) as a string and returns the result.
--Write a program with a closure to count the number of times a function is called.
*/

/*
## Task for Better Understanding:
--Create a function findMax that takes an array of integers as input and returns the largest number.
--Write a function reverseString that takes a string and returns the reversed string.
--Build a calculator program with a menu for different operations (add, subtract, multiply, divide). Use functions for each operation.
*/

func FindMax(arr []int) int {
	if len(arr) == 0 {
		panic("empty slice: cannot find maximum")
	}
	max := math.MinInt
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func ReversedString(s string) string {
	reverseString := ""
	for idx := range s {
		reverseString = reverseString + string(s[len(s)-1-idx])
	}
	return reverseString
}

// func main() {
// 	fmt.Println(ReversedString("akash"))
// }
