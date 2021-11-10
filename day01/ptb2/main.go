package main

import (
	"day01/readnumberfromkeyboard"
	"fmt"
	"math"
)

func main() {
	var a, b, c, x, x1, x2 float64
	var err error
	var Delta float64

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

	Delta = b*b - 4*a*c
	if Delta == 0 {
		x = -b / (2 * a)
		fmt.Println("Phuong trinh co nghiem kep x1 = x2 = ", x)
	}
	if Delta > 0 {
		x1 = (-b + math.Sqrt(Delta)) / (2 * a)
		x2 = (-b - math.Sqrt(Delta)) / (2 * a)
		fmt.Println("Phuong trinh co 2 nghiem phan biet: \n ")
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
	if Delta < 0 {
		fmt.Println("Phuong trinh vo nghiem")
	}
}
