package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func JSONFormat() {
	myJson := `
[
	{
		"first_name": "Akash",
		"last_name": "Kumar",
		"hair_color": "black",
		"has_dog": false
	},
	{
		"first_name": "gaurav",
		"last_name": "kumar",
		"hair_color": "brown",
		"has_dog": true
	}
]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshal json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person
	var m1 Person
	m1.FirstName = "captain"
	m1.LastName = "america"
	m1.HairColor = "blonde"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error marshal", err)
	}
	fmt.Println(string(newJson))
}
