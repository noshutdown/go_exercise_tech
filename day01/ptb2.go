package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readNumberFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	if result, err = strconv.ParseFloat(str, 64); err != nil {
		fmt.Println(err.Error())
		return 0.0, err
	}
	return result, nil
}

func main() {
	var a, b, c float64
	var err error
	var Delta float64

	fmt.Println("Nhap 3 so cua phuong trinh:")
	a, err = readNumberFromKeyboard("a: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	b, err = readNumberFromKeyboard("b: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	c, err = readNumberFromKeyboard("c: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	Delta = b*b - 4*a*c
	if Delta == 0 {
		var x float64
		x = -b / (2 * a)
		fmt.Println("Phuong trinh co nghiem kep x1 = x2 = ", x)
	}
	if Delta > 0 {
		var x1, x2 float64
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
