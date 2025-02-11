package main

import (
	"fmt"
	// "log"
	"strings"
)

func removeOccurrences(s string, part string) string {
	var i int
	l := len(part)

	for {
		i = strings.Index(s, part)
		if i < 0 {
			break
		}
		s = s[:i] + s[i+l:]
	}
	return s
}

func main() {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
	fmt.Println(removeOccurrences("axxxxyyyyb", "xy"))
	fmt.Println(removeOccurrences("eemckxmckx", "emckx"))
}
