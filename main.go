package main

import (
	"fmt"
	"time"
)

type Order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time // nanasecond precision
}

func newOrder(id string, amount float32, status string) *Order {
	order := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &order
}

// receiver type
func (o *Order) changeStatus(status string) {
	o.status = status
}

func (o *Order) getAmount() float32 {
	return o.amount
}

func main() {
	/*
		order := Order{
			id: "1",
			// amount: 23.232,
			status: "pending",
		}
		order.createAt = time.Now()
		order.changeStatus("delivered")
		fmt.Println(order.amount)
		//get status
		// fmt.Println(order.status)
	*/

	order := newOrder("1", 200, "pending")
	order.createAt = time.Now()
	// fmt.Println(order)
	fmt.Println(*order)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
}
