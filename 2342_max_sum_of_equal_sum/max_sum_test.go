package main

import (
	"fmt"
	"testing"

	"container/heap"
)

type item struct {
	v int
	s int
}
type IntHeap []item

func (i IntHeap) Len() int { return len(i) }
func (i IntHeap) Less(k, j int) bool {
	if i[k].s == i[j].s {
		return i[k].v > i[j].v
	}
	return i[k].s > i[j].s
}
func (i IntHeap) Swap(k, j int) { i[k], i[j] = i[j], i[k] }
func (i *IntHeap) Pop() any {
	old := *i
	n := len(old)
	v := old[n-1]
	*i = old[:n-1]
	return v
}
func (i *IntHeap) Push(x any) {
	*i = append(*i, x.(item))
}

func maximumSum(nums []int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		s := sumDigits(num)
		heap.Push(h, item{v: num, s: s})
	}

	set := false
	var last item
	max := -1
	for range h.Len() {
		current := heap.Pop(h).(item)
		// fmt.Println(current)
		if !set {
			last = current
			set = true
			continue
		}
		if current.s == last.s {
			s := current.v + last.v
			if s > max {
				max = s
			}
		}
		last = current
	}
	return max
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

func Test(t *testing.T) {
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
			got := maximumSum(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
