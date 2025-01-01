package main

import (
	"fmt"

	"main.go/helpers"
	"main.go/json"
)

const numPool = 100

func CalculateNumber(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// First()
	// VariableZeroValueBlankIdentifier()
	// TypeConversionScopeHousekeeping()
	// PrintfForDecimalHexadecimal()
	// Pointers()
	// TypeStruct()
	// ReceiverStructGo()
	// MapSlices()
	// DecisionMaking()
	// Loops()
	// Interface()
	// var myVar helpers.SomeType
	// myVar.TypeName = "Package"
	// myVar.TypeNumber = 2
	// fmt.Println(myVar)

	// Channel
	intChan := make(chan int)
	defer close(intChan)
	go CalculateNumber(intChan)
	num := <-intChan
	fmt.Println(num)

	// Json-formatter
	json.JSONFormat()
}
