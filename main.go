package main

import "fmt"

func main() {
	age := 14
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not an Adult")
	}

	budget := 2000
	if budget >= 5000 {
		fmt.Println("Go to Trip")
	} else if budget >= 1000 && budget < 5000 {
		fmt.Println("Go for Shopping")
	} else {
		fmt.Println("Stay home")
	}

	// declare variable in if construct
	if age := 12; age >= 18 {
		fmt.Println("Adult")
	} else if age >= 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("not an adult")
	}

}
