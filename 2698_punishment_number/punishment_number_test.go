package main

import (
	"fmt"
	"testing"
)

func punishmentNumber(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if canPartition(i*i, i) {
			count += i * i
		}
	}
	return count
}

func canPartition(num int, target int) bool {
	if target == num {
		return true
	}
	if target > num || target < 0 {
		return false
	}

	return canPartition(num/10, target-num%10) ||
		canPartition(num/100, target-num%100) ||
		canPartition(num/1000, target-num%1000)
}

func Test(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{10, 182},
		{37, 1478},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := punishmentNumber(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
