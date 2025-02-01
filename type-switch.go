package main

import "fmt"

// type-switch -> A type switch is a construct that allows you to compare the type of an interface against multiple types.

func TypeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an Int:: %d\n", v)
	case string:
		fmt.Printf("It's a string:: %s\n", v)
	case float32, float64:
		fmt.Printf("It's a floating value:: %v\n", v)
	case bool:
		fmt.Printf("It's a boolean value:: %v\n", v)
	default:
		fmt.Printf("Unknown Type %T\n", v)
	}
}
