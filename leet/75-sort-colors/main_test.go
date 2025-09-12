package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{2, 0, 2, 1, 1, 0},
			expect: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums:   []int{2, 0, 1},
			expect: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			sortColors(tt.nums)
			diff := cmp.Diff(tt.nums, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, tt.nums)
			}
		})
	}
}

func sortColors(nums []int) {
	l := len(nums)
	for i := range l {
		for j := range l - 1 - i {
			if nums[j] >= nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

