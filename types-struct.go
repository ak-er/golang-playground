package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	address   string
	age       int
	birthDate time.Time
}

func TypeStruct() {
	user := User{
		firstName: "Akash",
		lastName:  "Kumar",
	}
	fmt.Println(user.firstName, user.lastName, user.address, user.age, user.birthDate)
}
