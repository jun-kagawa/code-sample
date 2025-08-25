package main_test

import (
	"testing"
)

func TestCanJump(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{nums: []int{2, 3, 1, 1, 4}, expect: true},
		{nums: []int{3, 2, 1, 0, 4}, expect: false},
		{nums: []int{2, 0}, expect: true},
		{nums: []int{2, 5, 0, 0}, expect: true},
		{nums: []int{3, 0, 8, 2, 0, 0, 1}, expect: true},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := canJump(tt.nums); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func canJump(nums []int) bool {
	i := 0
	last := len(nums) - 1
	for i != last {
		if nums[i] == 0 {
			return false
		}
		if nums[i] == 1 {
			i++
			continue
		}
		var m int
		var s int
		for j := range nums[i] {
			if i+1+j > last {
				break
			}
			if nums[i+1+j]+j >= s {
				s = nums[i+1+j] + j
				m = i + 1 + j
			}
		}
		i = m
	}
	return true
}
