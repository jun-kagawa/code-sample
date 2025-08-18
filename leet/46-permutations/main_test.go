package main_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		nums   []int
		expect [][]int
	}{
		{nums: []int{1, 2, 3}, expect: [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
		{nums: []int{0, 1}, expect: [][]int{
			{0, 1},
			{1, 0},
		}},
		{nums: []int{1}, expect: [][]int{
			{1},
		}},
		{nums: []int{1,2,3,4}, expect: [][]int{
			{1},
		}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := permute(tt.nums)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}

}

func permute1(nums []int) [][]int {
	r := make([][]int, 0, len(nums))
	slices.Sort(nums)
	rr := make([]int, len(nums))
	copy(rr, nums)
	r = append(r, rr)
	l1 := len(nums) - 1
	for {
		i := l1
		for i > 0 && nums[i-1] > nums[i] {
			i--
		}
		if i <= 0 {
			break
		}
		j := l1
		for j >= i && nums[i-1] > nums[j] {
			j--
		}
		if j <= 0 {
			break
		}
		nums[i-1], nums[j] = nums[j], nums[i-1]
		s := nums[i:]
		slices.Sort(s)
		nums = append(nums[:i], s...)
		rr := make([]int, len(nums))
		copy(rr, nums)
		r = append(r, rr)
	}
	fmt.Println(r)
	return r
}

func permute(nums []int) [][]int {
	r := make([][]int, 0, len(nums))
	track(nums, 0, &r)
	fmt.Println(r)
	return r
}

func track(nums []int, start int, r *[][]int) {
	if start == len(nums) {
		rr := make([]int, len(nums))
		copy(rr, nums)
		*r = append(*r, rr)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]

		track(nums, start+1, r)

		nums[start], nums[i] = nums[i], nums[start]
	}
}

