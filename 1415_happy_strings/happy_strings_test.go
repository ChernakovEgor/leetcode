package main

import (
	"fmt"
	"testing"
)

func getHappyString(n int, k int) string {
	al := []string{"a", "b", "c"}
	res := ""
	count := 0

	var backtrack func(int) bool
	backtrack = func(currentIdx int) bool {
		if currentIdx == n {
			count++
			return count == k
		}
		for _, l := range al {
			if currentIdx > 0 && string(res[len(res)-1]) == l {
				continue
			}

			res += l
			if backtrack(currentIdx + 1) {
				return true
			}
			res = res[:currentIdx]
		}

		return false
	}

	backtrack(0)

	return res
}

func Test(t *testing.T) {
	cases := []struct {
		n    int
		k    int
		want string
	}{
		{1, 3, "c"},
		{1, 4, ""},
		{3, 9, "cab"},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := getHappyString(test.n, test.k)
			if got != test.want {
				t.Errorf("got %s want %s", got, test.want)
			}
		})
	}
}
