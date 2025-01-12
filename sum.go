package main

import "fmt"

func sum() {
	var a, b int
	fmt.Print("Enter a and b: ")
	fmt.Scanln(&a, &b)
	// fmt.Print("Enter b: ")
	// fmt.Scanln(&b)
	fmt.Printf("sum of %d + %d = %d\n", a, b, a+b)
}

/* ====  To Run this program uncomment below code (main block) ==== */
// func main() {
// 	sum()
// }
