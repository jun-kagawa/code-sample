package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		expect []int
	}{
		{
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := spiralOrder(tt.matrix)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}

func spiralOrder(matrix [][]int) []int {
	nextDirection := map[string]string{
		"top":    "right",
		"right":  "bottom",
		"bottom": "left",
		"left":   "top",
	}
	i, j := 0, 0
	x, y := len(matrix[0]), len(matrix)
	l := x * y
	r := make([]int, 0, x*y)
	direction := "top"

	for len(r) != l {
		if direction == "top" {
			for k := i; k < x; k++ {
				r = append(r, matrix[j][k])
			}
			j++
		} else if direction == "right" {
			for k := j; k < y; k++ {
				r = append(r, matrix[k][x-1])
			}
			x--
		} else if direction == "bottom" {
			for k := x - 1; k >= i; k-- {
				r = append(r, matrix[y-1][k])
			}
			y--
		} else {
			for k := y - 1; k >= j; k-- {
				r = append(r, matrix[k][i])
			}
			i++
		}
		direction = nextDirection[direction]
	}
	return r
}
