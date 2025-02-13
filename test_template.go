package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		input []int
		want  int
	}{}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := func(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
