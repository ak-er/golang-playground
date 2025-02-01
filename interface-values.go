package main

import (
	"fmt"
	"reflect"
)

/*

Interface Values -> An interface value consists of two components.
- Dynamic Type: The concrete type of the value.
- Dynamic Value: The actual value stored in the interface.

*/

func Describe(i interface{}) {
	fmt.Printf("Type: %v and value %v\n", reflect.TypeOf(i), i)
}

/*

Keypoints
- The dynamic type and value of interface can change at runtime.
- Use reflections or types assertions to inspect the dynamic type and value.

*/
