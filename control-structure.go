package main

import "fmt"

func ifElseStatement() {
	fmt.Println("if-else-statement")
	/*
		Syntax
		if condition {
			// if body
		}else if anotherCondition{
			// else if body
		}else {
			// else body
		}
	*/

	// example
	age := 100
	if age < 18 && age >= 10 {
		fmt.Println("minor")
	} else if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("children")
	}
}

func switchStatement() {
	n := 20
	switch n {
	case 1:
		fmt.Println("Hola")
	case 2:
		fmt.Println("Hi")
	case 3:
		fmt.Println("Hello")
	case 4:
		fmt.Println("Bonjure")
	default:
		fmt.Println("Namaste")
	}
}

// practice switch statement print(weekend, and weekdays)
func weekendAndweekday() {
	day := "january"
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("It's weekday")
	case "saturday", "sunday":
		fmt.Println("It's weekend")
	default:
		fmt.Println("Invalid day")

	}
}

// for loop

/*
-- Syntax
for initialization, condition, increment {
	// code to execute
}
*/

// print n numbers
func printN_Numbers() {
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

// also used for as while-loop
func printN_Numbers2() {
	n := 10
	i := 1
	for i <= n {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

/*
# break-continue statement
-- break: Exits the loop immediately.
-- continue: Skips the current iteration and moves to next one.
*/

func breakStatement() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func continueStatement() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

/*
## Exercises:
--Write a program to check if a number is positive, negative, or zero using if statements.
--Create a switch statement to determine the type of a vehicle based on the input (car, bike, truck, etc.).
--Use a for loop to print the first 10 natural numbers.
--Write a program that prints only odd numbers from 1 to 20 using for and continue.
--Create a program that exits a loop when the user enters a specific number (e.g., -1).
*/

/*
## Task for Better Understanding:
--Create a program that determines whether a year is a leap year using if statements.
--Write a program using a for loop to calculate the factorial of a given number.
--Use a switch statement to display the grade based on a student's score.
*/

// func main() {
// 	continueStatement()
// }
