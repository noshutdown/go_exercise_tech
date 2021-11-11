package ptb2

import (
	"fmt"
	"math"
)

func SolveEquation2(a, b, c float64) {

	Delta := b*b - 4*a*c
	if Delta == 0 {
		x := -b / (2 * a)
		fmt.Println("Phuong trinh co nghiem kep x1 = x2 = ", x)
	}
	if Delta > 0 {
		x1 := (-b + math.Sqrt(Delta)) / (2 * a)
		x2 := (-b - math.Sqrt(Delta)) / (2 * a)
		fmt.Println("Phuong trinh co 2 nghiem phan biet:")
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
	if Delta < 0 {
		fmt.Println("Phuong trinh vo nghiem")
	}
}
