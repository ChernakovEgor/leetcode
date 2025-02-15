package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func punishmentNumberSlow(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if checkPartitions(i) {
			count += i * i
		}
	}
	return count
}

func checkPartitions(num int) bool {
	numStr := strconv.Itoa(num * num)
	res := []string{}
	p(&res, numStr, 1)
	for _, n := range res {
		fields := strings.Fields(n)
		sum := 0
		for _, f := range fields {
			d, _ := strconv.Atoi(f)
			sum += d
		}
		if sum == num {
			return true
		}
	}
	return false
}

func p(res *[]string, num string, cur int) {
	if cur >= len(num) {
		*res = append(*res, num)
		return
	}

	v1 := num[:cur] + num[cur:]
	v2 := num[:cur] + " " + num[cur:]
	p(res, v1, cur+1)
	p(res, v2, cur+2)
}

func TestSlow(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{10, 182},
		{37, 1478},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			checkPartitions(test.want)
			got := punishmentNumberSlow(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
