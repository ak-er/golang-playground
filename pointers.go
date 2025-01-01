package main

import "fmt"

func Pointers() {
	var orgString string = "Hello! Python"
	fmt.Println("Orginal String: ", orgString)
	changeString(&orgString)
	fmt.Println("After changeString Orginal String: ", orgString)
}

func changeString(str *string) {
	fmt.Println("s", str)
	newString := "Hello! Go"
	*str = newString
}
