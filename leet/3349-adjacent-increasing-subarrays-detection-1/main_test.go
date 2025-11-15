package main_test

import (
	"testing"
)

func TestHasIncreasingSubArrays(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect bool
	}{
		{
			nums:   []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
			k:      3,
			expect: true,
		},
		{
			nums:   []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
			k:      5,
			expect: false,
		},
		{
			nums:   []int{-3, -19, -8, -16},
			k:      2,
			expect: false,
		},
		{
			nums:   []int{-15, 19},
			k:      1,
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := hasIncreasingSubarrays(tt.nums, tt.k); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func hasIncreasingSubarrays(nums []int, k int) bool {
	l := len(nums)
	for i := range l {
		if i+k > l {
			continue;
		}
		a := nums[i : i+k]
		if !isSorted(a) {
			continue
		}
		if i+k+k > l {
			continue;
		}
		b := nums[i+k : i+k+k]
		if isSorted(b) {
			return true
		}
	}
	return false
}

func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}
