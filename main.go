package main

import "fmt"

// for -> only construct of looping
func main() {
	// use while loop using for loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// for loop complete syntax
	/*
		for initilization; condition; increament/decreament {
			// excutable code
		}
	*/

	for j := 1; j <= 5; j++ {
		if j == 3 {
			continue
		}
		fmt.Println(j)
	}
	fmt.Println("================")
	// golang 1.22 version = range function
	for k := range 3 {
		fmt.Println(k)
	}
}
