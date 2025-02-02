package main

type Numericer interface {
	~int | ~float64
}

func Sum[T Numericer](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
