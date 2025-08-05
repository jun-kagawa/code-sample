package main_test

import (
	"fmt"
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		/*
		 */
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expect: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, expect: -1},
		{nums: []int{1}, target: 0, expect: -1},
		{nums: []int{1}, target: 1, expect: 0},
		{nums: []int{1, 2, 4, 5, 6, 7, 0}, target: 0, expect: 6},
		{nums: []int{1, 3}, target: 1, expect: 0},
		{nums: []int{2, 1}, target: 1, expect: 1},
		{nums: []int{3, 1}, target: 3, expect: 0},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := search(tt.nums, tt.target)
			if r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	i := 0
	j := len(nums) - 1
	var k int
	for i < j {
		if nums[i] > nums[i+1] {
			k = i + 1
			break
		}
		if nums[j] < nums[j-1] {
			k = j
			break
		}
		i++
		j--
	}
	if k != 0 {
		nums = append(nums[k:], nums[:k]...)
	}

	fmt.Println("k", k)
	fmt.Println("nums", nums)
	if idx, ok := slices.BinarySearch(nums, target); ok {
		r := idx + k
		if r >= len(nums) {
			r -= len(nums)
		}
		return r
	} else {
		return -1
	}
}
