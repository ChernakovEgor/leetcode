package main

import (
	"fmt"
)

func canMakeSubsequence(str1 string, str2 string) bool {
	s1 := []rune(str1)
	s2 := []rune(str2)
	i, j := 0, 0
	for j < len(s2) && i < len(s1) {
		if s1[i] == s2[j] || s2[j] == ('a'+(s1[i]+1-'a')%26) {
			i++
			j++
		} else {
			i++
		}
	}
	if j == len(s2) {
		return true
	}
	return false
}

func main() {
	fmt.Println(canMakeSubsequence("ab", "d"))
}
