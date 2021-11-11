package randomnumber

import "math/rand"

// Ham genarate number
func RandomNumber(min, max int) int {
	return rand.Intn(max-min) + min

}
