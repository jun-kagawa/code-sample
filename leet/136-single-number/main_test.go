package main_test

import (
	"testing"
)

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int, 0)
	for _, n := range nums {
		if m[n] == 0 {
			m[n] += 1
		} else {
			delete(m, n)
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func TestSingleNumber(t *testing.T) {
	if r := singleNumber([]int{2,2,1}); r != 1 {
		t.Errorf("expect: %d, result: %d", 1, r)
	}
	if r := singleNumber([]int{4,2,2,1,1}); r != 4 {
		t.Errorf("expect: %d, result: %d", 4, r)
	}
	if r := singleNumber([]int{1}); r != 1 {
		t.Errorf("expect: %d, result: %d", 1, r)
	}
}

