package main

import (
	ptb2 "day01/ptb2/solveEquation"
	"day01/readnumberfromkeyboard"
	"fmt"
)

func main() {
	var a, b, c float64
	var err error

	fmt.Println("Nhap 3 so cua phuong trinh:")
	a, err = readnumberfromkeyboard.ReadNumberFromKeyboard("a: ")
	if err != nil {
		return
	}
	b, err = readnumberfromkeyboard.ReadNumberFromKeyboard("b: ")
	if err != nil {
		return
	}
	c, err = readnumberfromkeyboard.ReadNumberFromKeyboard("c: ")
	if err != nil {
		return
	}

	ptb2.SolveEquation2(a, b, c)
}
