package main

import (
	"fmt"
)

// Basic of Interface
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says Woof", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says Meow", c.Name)
}

func main() {
	/*
		// Basic Interface
		dog := Dog{Name: "Buddy"}
		cat := Cat{Name: "Whiskers"}

		var speaker Speaker
		speaker = dog
		fmt.Println(speaker.Speak())
		speaker = cat
		fmt.Println(speaker.Speak())
	*/

	/*
		// empty interface
		PrintAnything(65)
		PrintAnything("this is the string")
		PrintAnything(23.231)
		PrintAnything(true)
	*/

	/*
		// Type Assertion
		TypeAssertion()
	*/

	/*
		// Type Switch
		type User struct{}
		TypeSwitch(23)
		TypeSwitch("golang")
		TypeSwitch(23.232)
		TypeSwitch(true)
		TypeSwitch(User{})
	*/

	/*
		// Embedding inteface
		EmbeddingInterface()
	*/

	// Interface values
	var i interface{}
	i = 23
	Describe(i)
	i = "Hello"
	Describe(i)
}
