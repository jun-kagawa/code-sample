package main_test

import (
	"testing"
)

func TestJump(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		/*
		 */
		{nums: []int{2, 3, 1, 1, 4}, expect: 2},
		{nums: []int{2, 3, 0, 1, 4}, expect: 2},
		{nums: []int{0}, expect: 0},
		{nums: []int{1, 2}, expect: 1},
		{nums: []int{2, 1}, expect: 1},
		{nums: []int{3, 2, 1}, expect: 1},
		{nums: []int{2, 3, 1}, expect: 1},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := jump(tt.nums); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func jump(nums []int) int {
	var k int
	l1 := len(nums)
	for i := 0; i < l1-1; {
		if i+nums[i] >= l1-1 {
			i += nums[i]
		} else {
			a := min(l1, i+nums[i]+1)
			j := maxx(nums[i+1 : a])
			i += j + 1
		}
		k++
	}
	return k
}

func maxx(nums []int) int {
	var index int
	var m int
	for i, n := range nums {
		if m <= (i + n) {
			m = i + n
			index = i
		}
	}
	return index
}
