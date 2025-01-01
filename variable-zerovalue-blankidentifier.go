package main

import "fmt"

func VariableZeroValueBlankIdentifier() {
	/*
		var a int = 23
		fmt.Println(a)
		b := "yes"
		fmt.Println(b)
		b = "hello"
		fmt.Println(b)
		var c = 34
		fmt.Println(c)
		x, y, z := "cosmon", 34, "galaxy"
		print(x, y, z)
		var g bool
		fmt.Println(g)
	*/
	a, b := 5, 10
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
