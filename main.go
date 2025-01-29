package main

import "fmt"

// iterating over data structure
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(slice)
	// iterate using for loop
	/*
		for i := 0; i < len(slice); i++ {
			fmt.Println(slice[i])
		}
	*/

	// iterate slice using range
	sum := 0
	for _, val := range slice {
		// fmt.Println(val)
		sum += val
	}
	fmt.Println(sum)

	// iterate map using range
	m := map[string]string{"fname": "Akash", "lname": "kumar"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// iterate string using range - unicode value for character
	str := "😊learning golang"
	for i, c := range str {
		fmt.Println(i, string(c))
	}
}
