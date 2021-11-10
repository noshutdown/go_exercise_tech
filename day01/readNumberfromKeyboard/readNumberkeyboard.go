package readnumberfromkeyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumberFromKeyboard(msg string) (result float64, err error) {
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
