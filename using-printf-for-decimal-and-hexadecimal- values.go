package main

import "fmt"

func PrintfForDecimalHexadecimal() {
	val := 45
	fmt.Printf("%v is in binary %b and hexadecimal %x\n", val, val, val)
	fmt.Printf("%v is in binary %b and hexadecimal %#x\n", val, val, val)
}
