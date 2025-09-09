package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		matrix [][]int
		expect [][]int
	}{
		{
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			expect: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			expect: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			setZeroes(tt.matrix)
			diff := cmp.Diff(tt.matrix, tt.expect)
			if diff != "" {
				t.Errorf("expect %v, actual: %v", tt.expect, tt.matrix)
			}
		})
	}
}

func setZeroes(matrix [][]int) {
	r, c := make(map[int]bool), make(map[int]bool)
	lm := len(matrix)
	ln := len(matrix[0])
	for m := range lm {
		for n := range ln {
			if matrix[m][n] == 0 {
				r[m] = true
				c[n] = true
			}
		}
	}
	for rr := range r {
		for n := range ln {
			matrix[rr][n] = 0
		}
	}
	for cc := range c {
		for m := range lm {
			matrix[m][cc] = 0
		}
	}
}
