package main

import "fmt"

func main() {
	// create an array
	var arr [4]int
	// assign value
	arr[2] = 34
	fmt.Println(arr)

	// length of array
	fmt.Println(len(arr))

	// another method to create an array
	arrr := [4]int{1, 2, 3, 4}
	fmt.Println(arrr)

	// 2d array
	ararr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(ararr)
}
