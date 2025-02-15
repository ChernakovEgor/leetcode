package main

import "fmt"

func main() {
	s := make([]int, 1, 5)
	fmt.Println(s)
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	s[0] = 10
	s = append(s, 1)
}
