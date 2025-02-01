package main

import "fmt"

// empty interface -> The emptry interface (interface{}) has no methods, so any type satisfies it. It is often used for functions that can accept any type of arguments.

func PrintAnything(val interface{}) {
	fmt.Println(val)
}
