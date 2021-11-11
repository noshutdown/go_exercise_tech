package main

import (
	"day01/random_number/comparenumber"
	"day01/random_number/randomnumber"
	"day01/readnumberfromkeyboard"
	"fmt"
)

func main() {
	var input float64
	var cInput int
	var err error
	input, err = readnumberfromkeyboard.ReadNumberFromKeyboard("ban hay doan 1 so: ")
	if err != nil {
		//		fmt.Print(err.Error())
		return
	}
	cInput = int(input)
	// Call randomnumber function
	number := randomnumber.RandomNumber(0, 100)
	fmt.Println("Random Number:", number)
	// call Comparenumber function
	comparenumber.RandomNumber(number, cInput)
}
