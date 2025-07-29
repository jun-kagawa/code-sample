package main_test

import (
	"slices"
)

func removeDuplicates(nums []int) int {
	/*
	limit := len(nums)
	m := make(map[int]struct{})
	for i := 0; i < limit; i++ {
		nn := nums[i]
		if _, ok := m[nn]; ok {
			t := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, t)
			i--
			limit--
		} else {
			m[nn] = struct{}{}
		}
	}
	return len(m)
	*/
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	arr := make([]int, 0, len(m))
	for k, _ := range m {
		arr = append(arr, k)
	}
	slices.Sort(arr)
	copy(nums, arr)
	return len(arr)
}

