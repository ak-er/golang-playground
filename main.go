package main

import "fmt"

func changeValue(num *int) {
	*num = 5
	fmt.Println("num value in changeValue func:", *num)
}

func main() {
	num := 3
	changeValue(&num)
	fmt.Println("num value after changeValue in main-func", num)
}
