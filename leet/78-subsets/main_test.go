package main_test

import (
	//	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums   []int
		expect [][]int
	}{
		{
			nums: []int{1, 2, 3},
			expect: [][]int{
				{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3},
			},
		},
		{
			nums: []int{0},
			expect: [][]int{
				{}, {0},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := subsets(tt.nums)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func subsets(nums []int) [][]int {
	l := len(nums)
	r := make([][]int, 0)
	back(nums, []int{}, &r, l)
	return r
}

func back(nums []int, cur []int, r *[][]int, l int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	(*r) = append((*r), tmp)

	if len(cur) == l {
		return
	}
	if len(nums) == 0 {
		return
	}

	for i := range nums {
		cur = append(cur, nums[i])
		t := make([]int, len(nums))
		copy(t, nums)
		back(t[i+1:], cur, r, l)
		cur = cur[:len(cur)-1]
	}
}
