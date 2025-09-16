package main_test

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect bool
	}{
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			expect: true,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := search(tt.nums, tt.target); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		}
		if nums[low] == nums[mid] {
			low++
			continue
		}
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}
