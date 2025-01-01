package helpers

import (
	"math/rand"
)

type SomeType struct {
	TypeName   string
	TypeNumber int
}

func RandomNumber(n int) int {
	// rand.Seed(time.Now().UnixNano()) // deprecated
	return rand.Intn(n)
}
