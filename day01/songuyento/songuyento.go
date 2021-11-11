package main

import (
	"day01/readnumberfromkeyboard"
	"fmt"
)

func main() {
	var i int
	// n = 30
	var err error
	nf, err := readnumberfromkeyboard.ReadNumberFromKeyboard("ban hay nhap 1 so de tim so nguyen to trong day so do: ")
	if err != nil {
		//		fmt.Print(err.Error())
		return
	}
	n := int(nf)
	songuyento := []int{}
	for n2 := 0; n2 <= n; n2++ {
		for i = 2; i < n2; i++ {
			if n2%i == 0 {
				break
			} else if i == n2-1 {
				songuyento = append(songuyento, n2)
			}
		}
	}
	fmt.Println(songuyento)
}
