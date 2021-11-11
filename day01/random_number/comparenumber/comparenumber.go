package comparenumber

import "fmt"

// Ham so sanh 2 so
func RandomNumber(randomnumber, inputnumber int) {
	switch {
	case inputnumber < randomnumber:
		fmt.Printf("So ban doan nho hon %d\n", randomnumber)
	case inputnumber > randomnumber:
		fmt.Printf("So ban doan lon hon %d\n", randomnumber)
	case inputnumber == randomnumber:
		fmt.Printf("Ban doan chinh xac!\n")
	}

}
