package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func twoSum(nums []int, target int) []int {
	return twoSum2(nums, target)
}
func twoSum1(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// x + y = target
// y = target - x

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		diff := target - num
		if j, ok := numMap[diff]; ok {
			return []int{i, j}
		}
		numMap[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			expect: []int{0, 1},
		},
		{
			nums:   []int{0, 4, 3, 0},
			target: 0,
			expect: []int{0, 3},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := twoSum(tt.nums, tt.target)
			if !cmp.Equal(r, tt.expect, cmpopts.SortSlices(func(i, j int) bool { return i < j })) {
				t.Errorf("test failed. %v", r)
			}
		})
	}
}
