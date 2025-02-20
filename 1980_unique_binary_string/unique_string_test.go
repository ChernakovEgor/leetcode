package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func findDifferentBinaryString(nums []string) string {
	seen := make(map[int64]bool)
	for _, num := range nums {
		r, _ := strconv.ParseInt(num, 2, 64)
		seen[r] = true
	}
	var i int64
	m := int64(math.Pow(float64(len(nums[0])), 2))
	for ; i <= m; i++ {
		if !seen[i] {
			res := strconv.FormatInt(i, 2)
			for len(res) < len(nums) {
				res = "0" + res
			}
			return res
		}
	}
	return ""
}

func Test(t *testing.T) {
	cases := []struct {
		input []string
		want  string
	}{
		{[]string{"01", "10"}, "00"},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := findDifferentBinaryString(test.input)
			if got != test.want {
				t.Errorf("got %s want %s", got, test.want)
			}
		})
	}
}
