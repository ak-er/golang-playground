package main

import "fmt"

func Loops() {
	// loop in go
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// loop in slice
	animals := []string{"horse", "deer", "elephant", "dog", "cat"}
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	// loop in maps
	mymap := make(map[string]string)
	mymap["1"] = "one"
	mymap["2"] = "two"
	mymap["3"] = "three"
	mymap["4"] = "four"
	mymap["5"] = "five"

	for key, value := range mymap {
		fmt.Println(key, value)
	}

	// loop in string
	mystring := "this is my string"
	for idx, letter := range mystring {
		fmt.Println(idx, letter)
		// fmt.Printf("%d %c\n", idx, letter)
	}
}
