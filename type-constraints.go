package main

import "golang.org/x/exp/constraints"

/*

Type Constraints: Type constraints restricts the types that can be used with generics. They ensures that the generics code works only with specific types.

Predefined constraints
- any: Allows any type.
- comparable: Allows types that can be compared using == and !=.
- constraints.Ordered: Allows types that support ordering (<, >, <=, >=).

*/

/*
	// Comparable type
	func printSlice[T comparable](slice []T) {
		for _, v := range slice {
			fmt.Println(v)
		}
	}
*/

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
