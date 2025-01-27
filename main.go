package main

import (
	"fmt"
	"slices"
)

func main() {
	/*
		Slices -> dynamic
		most used construct in go
		it haves more useful methods.
	*/
	// var nums = []int
	/*
		var nums = make([]int, 2, 5)
		fmt.Println(nums)
		fmt.Println("slice capacity:: ", cap(nums))
		fmt.Println("slice length:: ", len(nums))
		fmt.Println(nums == nil)
	*/

	var nums = make([]int, 0, 5)
	fmt.Println("slice capacity:: ", cap(nums))
	fmt.Println("slice length:: ", len(nums))
	// fmt.Println(nums == nil)

	nums = append(nums, 23)
	nums = append(nums, 76)
	nums = append(nums, 45)
	nums = append(nums, 34)
	nums = append(nums, 25)
	nums = append(nums, 46)
	nums = append(nums, 8)
	fmt.Println(nums)
	fmt.Println("slice capacity:: ", cap(nums))

	nums = []int{}
	nums = append(nums, 23)
	fmt.Println("slice capacity:: ", cap(nums))
	fmt.Println("slice length:: ", len(nums))
	// fmt.Println(nums)

	// copy functions
	var copyNums = make([]int, len(nums))
	nums = append(nums, 24)
	fmt.Println(nums)
	fmt.Println(copyNums)
	copy(copyNums, nums)
	fmt.Println(nums)
	fmt.Println(copyNums)

	// slice operator
	nums = []int{1, 2, 3, 4, 5}
	fmt.Println(nums[0:3])
	fmt.Println(nums[:3])
	fmt.Println(nums[3:])

	// slice
	nums = []int{1, 2}
	nums2 := []int{1, 3}
	fmt.Println(slices.Equal(nums, nums2))

	// 2d slices
	nums2D := [][]int{{1, 2}, {3, 4}}
	fmt.Println(nums2D)
}
