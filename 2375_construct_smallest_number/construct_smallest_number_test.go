package main

import (
	"fmt"
	"testing"
)

func smallestNumber(pattern string) string {
	permutations := make([]byte, len(pattern)+1)
	used := make([]bool, 10)

	var backtrack func(int) bool
	backtrack = func(currIdx int) bool {

		for i := 1; i <= 9; i++ {
			if used[i] {
				continue
			}

			if currIdx > 0 && pattern[currIdx-1] == 'I' {
				if permutations[currIdx-1] > byte(i)+'0' {
					continue
				}
			}

			if currIdx > 0 && pattern[currIdx-1] == 'D' {
				if permutations[currIdx-1] < byte(i)+'0' {
					continue
				}
			}

			permutations[currIdx] = byte(i) + '0'
			used[i] = true
			fmt.Println(string(permutations))

			if currIdx == len(pattern) {
				return true
			}

			if backtrack(currIdx + 1) {
				return true
			}

			used[i] = false
		}

		return false
	}

	backtrack(0)

	return string(permutations)
}

func Test(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"IIIDIDDD", "123549876"},
		{"DDD", "4321"},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := smallestNumber(test.input)
			if got != test.want {
				t.Errorf("got %s want %s", got, test.want)
			}
		})
	}
}
