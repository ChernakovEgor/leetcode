package main

import (
	"fmt"
	"strconv"
	"testing"
)

func smallestNumberStack(pattern string) string {
	stack := []int{}
	res := ""

	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, i+1)

		if i == len(pattern) || pattern[i] == 'I' {
			for j := len(stack) - 1; j >= 0; j-- {
				res += strconv.Itoa(stack[j])
				stack = stack[:j]
			}
		}
	}

	return res
}

func TestStack(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"IIIDIDDD", "123549876"},
		{"DDD", "4321"},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := smallestNumberStack(test.input)
			if got != test.want {
				t.Errorf("got %s want %s", got, test.want)
			}
		})
	}
}
