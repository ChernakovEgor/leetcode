package minoperationstoexceed

import (
	"container/heap"
	"fmt"
	"testing"
)

type pq []int

func (p pq) Len() int           { return len(p) }
func (p pq) Less(i, j int) bool { return p[i] < p[j] }
func (p pq) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(e any) {
	*p = append(*p, e.(int))
}
func (p *pq) Pop() any {
	old := *p
	v := old[len(old)-1]
	*p = old[:len(old)-1]
	return v
}

func minOperations(nums []int, k int) int {
	pq := &pq{}
	heap.Init(pq)
	for _, n := range nums {
		heap.Push(pq, n)
	}
	i := 0
	for ; pq.Len() >= 2; i++ {
		x := heap.Pop(pq).(int)
		if x >= k {
			return i
		}
		y := heap.Pop(pq).(int)
		heap.Push(pq, min(x, y)*2+max(x, y))
	}
	return i
}

func Test(t *testing.T) {
	cases := []struct {
		input []int
		k     int
		want  int
	}{
		{[]int{2, 11, 10, 1, 3}, 10, 2},
		{[]int{1, 1, 2, 4, 9}, 20, 4},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := minOperations(test.input, test.k)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
