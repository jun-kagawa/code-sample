package main_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct{
		nums []int
		expect [][]int
	}{
		{nums: []int{1,1,2}, expect: [][]int{
			{1,1,2},
			{1,2,1},
			{2,1,1},
		}},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := permuteUnique(tt.nums)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}

func permuteUnique(nums []int) [][]int {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		m[n]++
	}
	r := make([][]int, 0, len(nums))
	back(m, []int{}, &r, len(nums))
	fmt.Println(r)
	return r
}

func back(nums map[int]int, current []int, r *[][]int, l int) {
	if len(current) == l {
		rr := make([]int, l)
		copy(rr, current)
		*r = append(*r, rr)
		return
	}
	for k, v := range nums {
		if v == 0 {
			continue
		}
		nums[k]--
		current = append(current, k)
		back(nums, current, r, l)
		nums[k]++
		current = current[:len(current)-1]
	}
}
