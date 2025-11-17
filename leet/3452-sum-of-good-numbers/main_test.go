package main_test

import (
	"testing"
)

func TestSumOfGoodNumbers(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect int
	}{
		{
			nums:   []int{1, 3, 2, 1, 5, 4},
			k:      2,
			expect: 12,
		},
		{
			nums:   []int{2, 1},
			k:      1,
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := sumOfGoodNumbers(tt.nums, tt.k); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func sumOfGoodNumbers(nums []int, k int) int {
	var sum int
	for i, n := range nums {
		if i-k < 0 {
			if n > nums[i+k] {
				sum += n
			}
			continue
		}
		if i+k >= len(nums) {
			if n > nums[i-k] {
				sum += n
			}
			continue
		}
		if n > nums[i+k] && n > nums[i-k] {
			sum += n
		}
	}
	return sum
}
