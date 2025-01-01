package main

import (
	"fmt"
	"sort"
)

func MapSlices() {
	// don't recommend
	// var dontRecommendLikeMap map[string]string
	myMap := make(map[string]string)
	myMap["favoriteFood"] = "Pizza"

	fmt.Println(myMap)

	// length
	fmt.Println(len(myMap))
	// access
	fmt.Println(myMap["favoriteFood"])

	// slice
	var mySlice []string
	mySlice = append(mySlice, "world")
	mySlice = append(mySlice, "Hello")

	sort.Strings(mySlice)
	fmt.Println(mySlice)

	var numberSlice []int
	numberSlice = append(numberSlice, 4)
	numberSlice = append(numberSlice, 2)
	numberSlice = append(numberSlice, 1)
	numberSlice = append(numberSlice, 3)
	fmt.Println(numberSlice[0:2])
	sort.Ints(numberSlice)
	fmt.Println(numberSlice)

	// short hand
	s := []int{}
	s = append(s, 3)
	s = append(s, 1)
	fmt.Println(s)

}
