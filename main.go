package main

import "fmt"

// enumerated types

/*
type PaymentStatus int

const (
	Failed PaymentStatus = iota
	Pending
	Success
)
*/

type PaymentStatus string

const (
	Failed  PaymentStatus = "failed"
	Pending               = "pending"
	Success               = "success"
)

func checkPaymentStatus(status PaymentStatus) {
	fmt.Println("payment status:", status)
}

func main() {
	checkPaymentStatus(Pending)
}
