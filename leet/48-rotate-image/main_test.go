package main_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
		expect [][]int
	}{
		{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
			expect: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{matrix: [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
			expect: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{matrix: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
			expect: [][]int{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			rotate(tt.matrix)
			fmt.Println(tt.matrix)
			diff := cmp.Diff(tt.matrix, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}

func rotate(matrix [][]int) {
	r(matrix)
}

func r(m [][]int) {
	l := len(m)
	k := l - 1
	for i := 0; i < (l >> 1); i++ {
		for j := i; j < k-i; j++ {
			m[i][j], m[j][k-i], m[k-i][k-j], m[k-j][i] = m[k-j][i], m[i][j], m[j][k-i], m[k-i][k-j]
		}
	}
}
