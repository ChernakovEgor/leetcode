package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	i := 0
	l := len(s)
	for i < l {
		if unicode.IsDigit(rune(s[i])) {
			if i > 0 {
				s = s[:i-1] + s[i:]
				i--
				l--
			}
			s = s[:i] + s[i+1:]
			i--
			l--
		}
		i++
	}
	return s
}

func clearDigitsMem(s string) string {
	res := make([]byte, len(s))
	var skipped int
	pos := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			skipped++
		} else {
			if skipped == 0 {
				res[pos] = s[i]
				pos--
			} else {
				skipped--
			}
		}
	}

	return string(res[pos+1:])
}

func main() {
	fmt.Println(clearDigitsMem("abc"))
	fmt.Println(clearDigitsMem("abc1"))
	fmt.Println(clearDigitsMem("cb34"))
	fmt.Println(clearDigitsMem("1"))
}
