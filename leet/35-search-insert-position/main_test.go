package main_test

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		/*
		*/
		{nums: []int{1, 3, 5, 6}, target: 5, expect: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expect: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expect: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, expect: 0},
		{nums: []int{1, 3, 5, 6, 8}, target: 0, expect: 0},
		{nums: []int{1, 3, 5, 7, 9}, target: 10, expect: 5},
		{nums: []int{1, 3, 5, 7, 9}, target: 6, expect: 3},
		{nums: []int{1}, target: 6, expect: 1},
		{nums: []int{1}, target: 0, expect: 0},
		{nums: []int{1}, target: 1, expect: 0},
		{nums: []int{1, 3}, target: 3, expect: 1},
		{nums: []int{1, 3}, target: 1, expect: 0},
		{nums: []int{3, 4, 5, 6, 7, 8}, target: 6, expect: 3},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := searchInsert(tt.nums, tt.target)
			if r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	h := -1
	for start < end && h != 0 {
		h = ((end - start) >> 1)
		p := start + h
		fmt.Println(h)
		if nums[p] > target {
			end -= h
		} else if nums[p] < target {
			start += h
		} else if nums[p] == target {
			return h + start
		}
	}
	fmt.Println("start", start, "end", end)
	if nums[start] < target && nums[end] < target {
		return end + 1
	} else if  nums[start] < target && target <= nums[end] {
		return end
	} else {
		return start
	}
}
