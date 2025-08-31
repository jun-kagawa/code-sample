package main_test

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m      int
		n      int
		expect int
	}{
		{m: 1, n: 1, expect: 1},
		{m: 3, n: 1, expect: 1},
		{m: 3, n: 2, expect: 3},
		{m: 3, n: 3, expect: 6},
		{m: 3, n: 4, expect: 10},
		{m: 3, n: 5, expect: 15},
		{m: 3, n: 6, expect: 21},
		{m: 3, n: 7, expect: 28},
		{m: 4, n: 1, expect: 1},
		{m: 4, n: 2, expect: 4},
		{m: 4, n: 3, expect: 10},
		{m: 4, n: 4, expect: 20},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := uniquePaths(tt.m, tt.n); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

var mm int
var nn int

func uniquePaths1(m int, n int) int {
	c := 1
	for i := 1; i < m; i++ {
		c = c * (n - 1 + i) / i
	}
	return c
}

func uniquePaths(m int, n int) int {
	prev := make([]int, n)
	cur := make([]int, n)
	prev[0] = 1
	for range m {
		cur[0] = prev[0]
		for j := 1; j < n; j++ {
			cur[j] = cur[j-1] + prev[j]
		}
		prev, cur = cur, prev
	}
	return prev[n-1]
}

func back(x, y int, c *int) {
	if x == mm && y == nn {
		*c++
		return
	}
	if x <= mm {
		x++
		back(x, y, c)
		x--
	}

	if y <= nn {
		y++
		back(x, y, c)
		y--
	}
}
