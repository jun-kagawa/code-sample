package main_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{input: []int{1, 1, 5}, expect: []int{1, 5, 1}},
		{input: []int{1, 2, 3}, expect: []int{1, 3, 2}},
		{input: []int{1, 3, 2}, expect: []int{2, 1, 3}},
		{input: []int{3, 2, 1}, expect: []int{1, 2, 3}},
		{input: []int{2, 3, 1}, expect: []int{3, 1, 2}},
		{input: []int{2, 1, 3}, expect: []int{2, 3, 1}},
		{input: []int{2, 3, 1}, expect: []int{3, 1, 2}},
		{input: []int{3, 1, 2}, expect: []int{3, 2, 1}},
		{input: []int{1, 2, 3, 4}, expect: []int{1, 2, 4, 3}},
		{input: []int{1, 3, 2, 4}, expect: []int{1, 3, 4, 2}},
		{input: []int{1, 3, 4, 2}, expect: []int{1, 4, 2, 3}},
		{input: []int{1, 4, 2, 3}, expect: []int{1, 4, 3, 2}},
		{input: []int{2, 1, 3, 4}, expect: []int{2, 1, 4, 3}},
		{input: []int{2, 3, 1, 4}, expect: []int{2, 3, 4, 1}},
		{input: []int{2, 3, 4, 1}, expect: []int{2, 4, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("input: %v", tt.input), func(t *testing.T) {
			nextPermutation(tt.input)
			if diff := cmp.Diff(tt.input, tt.expect); diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	re(nums, i+1, n-1)
}

func re(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func nextPermutation1(nums []int) {
	var isSorted bool
	i := 0
	for i < len(nums) {
		for j := i; j < len(nums)-1; j++ {
			if nums[j] <= nums[j+1] {
				if j+1 == len(nums)-1 {
					isSorted = true
				}
				continue
			} else {
				break
			}
		}
		if isSorted {
			break
		}
		i++
	}
	fmt.Println("i", i)

	if i == 0 {
		nums[1], nums[2] = nums[2], nums[1]
		return
	}
	if i == len(nums) {
		prev := nums[0]
		slices.Sort(nums)
		var n int
		for k, num := range nums {
			if prev < num {
				n = k
				break
			}
		}
		a := []int{nums[n]}
		arr := nums[:n+copy(nums[n:], nums[n+1:])]
		fmt.Println("a", a, "arr", arr)
		a = append(a, arr...)
		copy(nums, a)
		return
	}
	nums[i], nums[i+1] = nums[i+1], nums[i]
}
