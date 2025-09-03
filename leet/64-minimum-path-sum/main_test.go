package main_test

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid   [][]int
		expect int
	}{
		{grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, expect: 7},
		{grid: [][]int{{1, 2, 3}, {4, 5, 6}}, expect: 12},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := minPathSum(tt.grid); r != tt.expect {
				fmt.Println(r)
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	pre := grid[0]
	for i := range pre {
		if i != 0 {
			pre[i] += pre[i-1]
		}
	}
	cur := make([]int, n)
	for i := 1; i < m; i++ {
		cur[0] = pre[0] + grid[i][0]
		for j := 1; j < n; j++ {
			cur[j] = min(cur[j-1], pre[j]) + grid[i][j]
		}
		pre, cur = cur, pre
	}
	return pre[n-1]
}
