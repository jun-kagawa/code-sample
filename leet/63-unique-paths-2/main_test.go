package main_test

import (
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		obstacleGrid [][]int
		expect       int
	}{
		{obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, expect: 2},
		{obstacleGrid: [][]int{{0, 0}, {0, 1}}, expect: 0},
	}

	for _, tt := range tests {
		t.Run("error", func(t *testing.T) {
			if r := uniquePathsWithObstacles(tt.obstacleGrid); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	prev := make([]int, n)
	curr := make([]int, n)
	prev[0] = 1
	for i := range m {
		if obstacleGrid[i][0] == 1 {
			curr[0] = 0
		} else {
			curr[0] = prev[0]
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				curr[j] = 0
			} else {
				curr[j] = curr[j-1] + prev[j]
			}
		}
		prev, curr = curr, prev
	}
	return prev[n-1]
}

