package main

import (
	"fmt"
	"testing"
)

func numTilePossibilities(tiles string) int {
	var backtrack func(int, string)
	permutations := ""
	seen := make(map[string]bool)
	backtrack = func(_ int, avl string) {
		for i := 0; i < len(avl); i++ {
			permutations += string(avl[i])
			// fmt.Printf("%s\t%s\t%s\n", avl, permutations, avl[:i]+avl[i+1:])
			seen[permutations] = true
			backtrack(1, avl[:i]+avl[i+1:])
			permutations = permutations[:len(permutations)-1]
		}
	}

	backtrack(0, tiles)

	return len(seen)
}

func Test(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"AAB", 8},
		{"AAABBC", 188},
		{"V", 1},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := numTilePossibilities(test.input)
			if got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
