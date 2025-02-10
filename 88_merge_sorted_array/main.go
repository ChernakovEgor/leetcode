package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func main() {
	n := []int{2, 0}
	m := []int{1}
	merge(n, 1, m, 1)
	fmt.Println(n)
}
