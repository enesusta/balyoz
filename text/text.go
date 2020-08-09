package text

import "unicode"

func CapitalizeWithTurkish(str string) string {
	tmp := []rune(str)
	tmp[0] = unicode.TurkishCase.ToUpper(tmp[0])
	return string(tmp)
}
