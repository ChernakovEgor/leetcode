package main

import (
	"fmt"
	"reflect"
	"testing"
)

func permute(nums []int) [][]int {
	res := [][]int{}
	used := make(map[int]bool)
	c := make([]int, len(nums))

	var p func(curIndex int)
	p = func(curIndex int) {
		if curIndex == len(nums) {
			r := make([]int, len(nums))
			copy(r, c)
			res = append(res, r)
			fmt.Println(c)
			return
		}
		for _, n := range nums {
			if !used[n] {
				c[curIndex] = n
				used[n] = true
				p(curIndex + 1)
				used[n] = false
			}
		}
	}

	p(0)

	return res
}

func Test(t *testing.T) {
	cases := []struct {
		input []int
		want  [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := permute(test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
