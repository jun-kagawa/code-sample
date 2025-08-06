package main_test

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, expect: []int{3, 4}},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 6, expect: []int{-1, -1}},
		{nums: []int{}, target: 0, expect: []int{-1, -1}},
		{nums: []int{1}, target: 1, expect: []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := searchRange(tt.nums, tt.target)
			if diff := cmp.Diff(r, tt.expect); diff != "" {
				t.Errorf("%v", diff)
			}
		})
	}
}

func searchRange(nums []int, target int) []int {
	i, ok := slices.BinarySearch(nums, target)
	if !ok {
		return []int{-1, -1}
	}
	j := i
	for ; j < len(nums) && nums[i] == nums[j]; j++ {
	}
	return []int{i, j-1}
}
