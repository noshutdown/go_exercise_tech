package main

import (
	"day01/readnumberfromkeyboard"
	"fmt"
)

// func checksonguyento(n int) {
// 	var i int
// 	for i = 2; i < (n - 1); i++ {
// 		if n%i == 0 {
// 			return
// 		}
// 	}
// 	fmt.Println(i)
// }
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
	songuyento := []int{1, 2}
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
