package main

import (
	"day01/readnumberfromkeyboard"
	"fmt"
	"math/rand"
)

func main() {
	var input float64
	var min, max, cInput int
	var err error
	input, err = readnumberfromkeyboard.ReadNumberFromKeyboard("ban hay doan 1 so: ")
	if err != nil {
		//		fmt.Print(err.Error())
		return
	}
	cInput = int(input)

	min = 0
	max = 100
	random := (rand.Intn(max-min) + min)
	fmt.Printf("So ngau nhien la: %d\n", random)

	switch {
	case cInput < random:
		fmt.Printf("So ban doan nho hon %d\n", random)
	case cInput > random:
		fmt.Printf("So ban doan lon hon %d\n", random)
	case cInput == random:
		fmt.Printf("Ban doan chinh xac!\n")
	}
}
