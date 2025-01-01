package main

import "fmt"

func DecisionMaking() {
	var age int
	age = 10

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Teen age")
	} else {
		fmt.Println("Children")
	}

	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("Enter valid number of day")
	}
}
