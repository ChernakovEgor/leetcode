package main

import (
	"fmt"
	"testing"
)

func maximumSumLinear(nums []int) int {
	digitSums := [82]int{}
	sum := -1
	for _, num := range nums {
		s := sumDigits(num)
		if digitSums[s] > 0 {
			if digitSums[s]+num > sum {
				sum = digitSums[s] + num
			}
		} else {
			digitSums[s] = num
		}
		digitSums[s] = max(digitSums[s], num)
	}
	return sum
}

func sumDigits(num int) (sum int) {
	for {
		if num == 0 {
			break
		}
		sum += num % 10
		num = num / 10
	}
	return sum
}

func TestLinear(t *testing.T) {
	cases := []struct {
		input []int
		want  int
	}{
		{[]int{18, 43, 36, 13, 7}, 54},
		{[]int{10, 12, 19, 14}, -1},
		{[]int{368, 369, 307, 304, 384, 138, 90, 279, 35, 396, 114, 328, 251, 364, 300, 191, 438, 467, 183}, 835},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := maximumSumLinear(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
