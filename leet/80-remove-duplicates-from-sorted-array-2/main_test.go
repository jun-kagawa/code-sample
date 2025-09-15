package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums   []int
		arr    []int
		expect int
	}{
		{
			nums:   []int{1, 1, 1, 2, 2, 3},
			arr:    []int{1, 1, 2, 2, 3},
			expect: 5,
		},
		{
			nums:   []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			arr:    []int{0, 0, 1, 1, 2, 3, 3},
			expect: 7,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := removeDuplicates(tt.nums)
			diff := cmp.Diff(tt.nums[:r], tt.arr[:r])
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.arr, tt.nums)
			}
			if r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	count, current := 0, 1
	for i := 1; i < l; i++ {
		if nums[i-1] != nums[i] {
			count = 0
			nums[current] = nums[i]
			current++
		} else {
			count++
			if count < 2 {
				nums[current] = nums[i]
				current++
			}
		}
	}
	return current
}

func removeDuplicates2(nums []int) int {
	var i int = 1
	var f bool
	for i < len(nums) {
		if nums[i-1] == nums[i] {
			if f {
				nums = nums[:i+copy(nums[i:], nums[i+1:])]
				continue
			} else {
				f = true
			}
		} else {
			f = false
		}
		i++
	}
	return len(nums)
}
