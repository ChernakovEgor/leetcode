package main

import (
	"fmt"
	"reflect"
	"testing"
)

func subsets(nums []int) [][]int {
	res := [][]int{}
	permutation := []int{}
	var backtrack func(int, int)
	backtrack = func(currentIdx int, startNum int) {
		e := make([]int, len(permutation))
		copy(e, permutation)
		res = append(res, e)

		for i := startNum; i < len(nums); i++ {
			permutation = append(permutation, nums[i])
			backtrack(currentIdx+1, i+1)
			permutation = permutation[:currentIdx]
		}
	}

	backtrack(0, 0)

	return res
}

func Test(t *testing.T) {
	cases := []struct {
		input []int
		want  [][]int
	}{
		{[]int{0}, [][]int{{}, {0}}},
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := subsets(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
