package main_test

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		args   []int
		expect [][]int
	}{
		{args: []int{-1, 0, 1, 2, -1, -4}, expect: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
		},
		{args: []int{0, 1, 1}, expect: [][]int{}},
		{args: []int{0, 0, 0}, expect: [][]int{
			{0, 0, 0},
		}},
		{args: []int{0, 0, 0, 0}, expect: [][]int{
			{0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := threeSum(tt.args)
			expect := tt.expect
			if len(expect) != len(r) {
				t.Errorf("missing same len. expect: %v, actual: %v", len(expect), len(r))
			}
			for i := range len(expect) {
				for j := range len(expect[i]) {
					if expect[i][j] != r[i][j] {
						t.Errorf("expect: %v, actual: %v", expect, r)
					}
				}
			}

		})
	}
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	r := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
	j := i + 1
	k := len(nums) - 1
	for j < k {
		sum := nums[i] + nums[j] + nums[k]
		if sum == 0 {
			s := []int{nums[i], nums[j], nums[k]}
			r = append(r, s)
			j++
			k--
			for j < k && nums[j] == nums[j-1] {
				j++
			}
			for j < k && nums[k] == nums[k+1] {
				k--
			}
		}
		if sum > 0 {
			k--
		}
		if sum < 0 {
			j++
		}
	}
	}
	return r
}

func threeSum1(nums []int) [][]int {
	set := make(map[string]bool)
	r := make([][]int, 0)
	slices.Sort(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			t := 0 - (nums[i] + nums[j])
			start := min(j+1, len(nums))
			ns := nums[start:]
			k, ok := slices.BinarySearch(ns, t)
			/*
				if nums[i] == -1 && nums[j] == -1 {
					fmt.Println(k, ok)
				}
			*/
			if ok {
				s := []int{nums[i], nums[j], ns[k]}
				slices.Sort(s)
				var sb strings.Builder
				for _, ss := range s {
					sb.WriteString(strconv.Itoa(ss))
				}
				if _, ok := set[sb.String()]; !ok {
					r = append(r, s)
					set[sb.String()] = true
				}
			}
		}
	}
	return r
}
