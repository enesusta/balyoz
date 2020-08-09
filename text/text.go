package text

import (
	"strings"
	"unicode"
)

func CapitalizeWithTurkish(str string) string {
	var builder strings.Builder
	first := string(str[0])
	remain := str[1:]
	first = strings.ToUpperSpecial(unicode.TurkishCase, first)

	builder.WriteString(first)
	builder.WriteString(remain)
	return builder.String()
}