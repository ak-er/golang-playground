package main

/*

Type Inference: Go can often infer the type arguments for generic functions, so you don't need to specify them explicitly.

*/

func Identity[T any](v T) T {
	return v
}
