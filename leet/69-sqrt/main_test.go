package main_test

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x      int
		expect int
	}{
		{x: 4, expect: 2},
		{x: 8, expect: 2},
		{x: 16, expect: 4},
		{x: 0, expect: 0},
		{x: 1, expect: 1},
		{x: 6, expect: 2},
		{x: 9, expect: 3},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := mySqrt(tt.x); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func mySqrt(x int) int {
	start, end := 0, x+1
	for start < end {
		mid := start + (end-start)/2
		if mid*mid > x {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start - 1
}
