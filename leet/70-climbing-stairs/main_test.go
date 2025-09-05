package main_test

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{n: 1, expect: 1},
		{n: 2, expect: 2},
		{n: 3, expect: 3},
		{n: 4, expect: 5},
		{n: 5, expect: 8},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := climbStairs(tt.n); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func climbStairs(n int) int {
	r, next, second := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			r = 1
		} else if i == 2 {
			r = 2
		} else {
			r = next + second
		}
		next, second = second, r
	}
	return r
}

