package main_test

import (
	"math"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expect: 6},
		{nums: []int{1}, expect: 1},
		{nums: []int{5, 4, -1, 7, 8}, expect: 23},
		{nums: []int{-1}, expect: -1},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := maxSubArray(tt.nums); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func maxSubArray(nums []int) int {
	var sum int = math.MinInt
	var cur int = 0
	for _, n  := range nums {
		cur += n
		if cur > sum {
			sum = cur
		}
		if cur < 0 {
			cur = 0
		}
	}
	return sum
}
