package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	result := sum(3, 2, 1, 4, 5, 6, 4)
	fmt.Println(result)

	// using slice
	s := []int{1, 2, 3, 4, 5}
	result = sum(s...)
	fmt.Println(result)
}
