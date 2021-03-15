package main

import (
	"fmt"
	"github.com/scyanh/Aspiration/mapper"
	"unicode"
)

const ASPIRATION = "Aspiration.com"

func main() {
	// capitalize with function
	fmt.Println(CapitalizeEveryThirdAlphanumericChar(ASPIRATION))

	// capitalize with mapper package
	s, err := mapper.NewSkipString(3, ASPIRATION)
	if err != nil {
		fmt.Println(err)
		return
	}
	mapper.MapString(&s)
	fmt.Println(&s)
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {
	frequency := 3
	offset := 0

	transformRune := func(pos int, runes []rune) string {
		runeValid := false
		if runes[pos] >= 97 && runes[pos] <= 122 ||
			runes[pos] >= 65 && runes[pos] <= 90 ||
			runes[pos] >= 48 && runes[pos] <= 57 {
			runeValid = true
		}
		if !runeValid {
			offset++
			return s
		}

		if frequency == 1 {
			runes[pos] = unicode.ToUpper(runes[pos])
		} else if pos > 0 && (pos+1-offset)%frequency == 0 {
			runes[pos] = unicode.ToUpper(runes[pos])
		} else {
			runes[pos] = unicode.ToLower(runes[pos])
		}

		return string(runes)
	}

	r := []rune(s)

	for pos, _ := range r {
		s = transformRune(pos, r)
	}

	return s
}
