package main

import (
	"day01/readnumberfromkeyboard"
	songuyento "day01/tim_songuyento/checksonguyento"
	"fmt"
)

func main() {
	// n = 30
	var err error
	nf, err := readnumberfromkeyboard.ReadNumberFromKeyboard("ban hay nhap 1 so de tim so nguyen to trong day so do: ")
	if err != nil {
		//		fmt.Print(err.Error())
		return
	}
	n := int(nf)

	fmt.Println(songuyento.CheckPrime(n))
}
