package main_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		n      int
		expect [][]int
	}{
		{n: 3, expect: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{n: 1, expect: [][]int{{1}}},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := generateMatrix(tt.n)
			fmt.Println("r", r)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}

func generateMatrix(n int) [][]int {
	r := make([][]int, 0, n)
	for range n {
		r = append(r, make([]int, n))
	}
	nextDirection := map[string]string{
		"top":    "right",
		"right":  "bottom",
		"bottom": "left",
		"left":   "top",
	}
	i, j := 0, 0
	x, y := n, n
	direction := "top"
	for k := 1; k < (n*n)+1; {
		if direction == "top" {
			for l := i; l < x; l++ {
				r[j][l] = k
				k++
			}
			j++
		} else if direction == "right" {
			for l := j; l < y; l++ {
				r[l][x-1] = k
				k++
			}
			x--
		} else if direction == "bottom" {
			for l := x - 1; l >= i; l-- {
				r[y-1][l] = k
				k++
			}
			y--
		} else {
			for l := y - 1; l >= j; l-- {
				r[l][i] = k
				k++
			}
			i++
		}
		direction = nextDirection[direction]
	}
	return r
}
