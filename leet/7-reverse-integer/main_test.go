package main_test

import (
	"math"
	"testing"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		res = res*10 + pop
		if res < math.MinInt32 || res > math.MaxInt32 {
			return 0
		}
	}
	return res
}

func TestReverse(t *testing.T) {
	tests := []struct {
		x      int
		expect int
	}{
		{x: 123, expect: 321},
		{x: -123, expect: -321},
		{x: 120, expect: 21},
		{x: 1534236469, expect: 0},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := reverse(tt.x)
			if r != tt.expect {
				t.Errorf("failed. expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}
