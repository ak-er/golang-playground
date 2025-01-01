package main

import "fmt"

func TypeConversionScopeHousekeeping() {
	a := 23
	fmt.Printf("Value of a is %v and the type of value is %T\n", a, a)
	var x = 23.98
	fmt.Printf("Value of x is %v and the type of value is %T\n", x, x)
	var f = float32(34.23)
	fmt.Printf("Value of f is %v and the type of value is %T\n", f, f)
	x = float64(f)
	fmt.Printf("Value of x is %.4f and the type of value is %T\n", x, x)
}
