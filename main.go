package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name  string
	phone string
}

func newCustomer(name string, phone string) *Customer {
	customer := Customer{
		name:  name,
		phone: phone,
	}
	return &customer
}

type Order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time // nanasecond precision
	customer Customer
}

func main() {
	order := Order{
		id:       "1",
		amount:   23.232,
		status:   "pending",
		customer: *newCustomer("john", "1234567890"),
	}
	order.createAt = time.Now()
	fmt.Println(order)
	fmt.Println(order.customer.name)

}
