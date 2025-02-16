package lexicographicallylargest

import (
	"fmt"
	"reflect"
	"testing"
)

func constructDistancedSequence(n int) []int {
	res := make([]int, 2*n-1)
	used := make(map[int]bool)

	var backtrack func(int) bool
	backtrack = func(currentIndex int) bool {
		if currentIndex == len(res) {
			return true
		}

		if res[currentIndex] != 0 {
			return backtrack(currentIndex + 1)
		}

		for i := n; i >= 1; i-- {
			if used[i] {
				continue
			}

			if i == 1 {
				res[currentIndex] = 1
				used[1] = true
				if backtrack(currentIndex + 1) {
					return true
				}

				//back
				res[currentIndex] = 0
				used[1] = false
			}

			if currentIndex+i >= len(res) || res[currentIndex+i] != 0 {
				continue
			}

			res[currentIndex] = i
			res[currentIndex+i] = i
			used[i] = true
			if backtrack(currentIndex + 1) {
				return true
			}

			res[currentIndex] = 0
			res[currentIndex+i] = 0
			used[i] = false
		}

		return false
	}

	backtrack(0)
	return res
}

func Test(t *testing.T) {
	cases := []struct {
		input int
		want  []int
	}{
		{3, []int{3, 1, 2, 3, 2}},
		{5, []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := constructDistancedSequence(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
