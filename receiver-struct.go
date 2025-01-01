package main

import "fmt"

type Customer struct {
	firstName string
	lastName  string
}

// Receivers struct with function
func (c *Customer) returnFullName() string {
	return c.firstName + " " + c.lastName
}

func ReceiverStructGo() {
	var customer Customer
	customer.firstName = "Akash"
	customer.lastName = "Kumar"

	newCustomer := Customer{
		firstName: "Akash",
		lastName:  "Kumar",
	}

	fmt.Println("customer fullName:", customer.firstName, customer.lastName)
	fmt.Println("new customer fullname:", newCustomer.firstName, newCustomer.lastName)

	fmt.Println("receiver:: customer fullname", customer.returnFullName())
	fmt.Println("receiver:: new customer fullname", newCustomer.returnFullName())

}
