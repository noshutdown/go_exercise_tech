package main

import (
	"day01/readnumberfromkeyboard"
	"fmt"
)

func main() {
	//fmt.Println("Nhap 1 so:")
	var n int
	var err error
	v, err := readnumberfromkeyboard.ReadNumberFromKeyboard("Nhap 1 so:")
	if err != nil {
		//		fmt.Print(err.Error())
		return
	}
	n = int(v)
	// m := n / 2
	if n < 2 {
		fmt.Printf("%d khong phai la so nguyen to\n", n)
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d khong phai la so nguyen to\n", n)
			break

		} else if i == n-1 {
			fmt.Printf("%d la so nguyen to\n", n)
		}
	}
}
