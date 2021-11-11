package main

import (
	checkprime "day01/check_songuyento/checksonguyento"
	"day01/readnumberfromkeyboard"
	"fmt"
)

func main() {

	var n int
	var err error
	v, err := readnumberfromkeyboard.ReadNumberFromKeyboard("Nhap 1 so: ")
	if err != nil {
		return
	}
	n = int(v)
	if checkprime.Checkprime(n) {
		fmt.Printf("%d la so nguyen to\n", n)
	} else {
		fmt.Printf("%d khong phai la so nguyen to\n", n)
	}
}
