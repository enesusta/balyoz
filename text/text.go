package text

import (
	"fmt"
	"strconv"
	"unicode"
)

func CapitalizeWithTurkish(str string) string {
	tmp := []rune(str)
	tmp[0] = unicode.TurkishCase.ToUpper(tmp[0])
	return string(tmp)
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	temp := 0

	if err != nil {
		fmt.Printf("InvalidNumberFormatException occured.")
		temp = 0
	}

	temp = i

	println(temp)

	return temp
}
