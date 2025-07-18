package main_test

import (
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	if r := threeSumClosest([]int{-1, 2, 1, -4}, 1); r != 2 {
		t.Errorf("expect: %v, actual: %v", 2, r)
	}
	if r := threeSumClosest([]int{0, 0, 0}, 1); r != 0 {
		t.Errorf("expect: %v, actual: %v", 0, r)
	}
	if r := threeSumClosest([]int{-4, -1, 1, 2, 4}, 3); r != 2 {
		t.Errorf("expect: %v, actual: %v", 4, r)
	}
	if r := threeSumClosest([]int{-4, -1, 1, 2, 4, 5}, 3); r != 3 {
		t.Errorf("expect: %v, actual: %v", 4, r)
	}
	if r := threeSumClosest([]int{1, 1, 1, 1}, 0); r != 3 {
		t.Errorf("expect: %v, actual: %v", 3, r)
	}
	if r := threeSumClosest([]int{4,0,5,-5,3,3,0,-4,-5}, -2); r != -2 {
		t.Errorf("expect: %v, actual: %v", -2, r)
	}
	if r := threeSumClosest([]int{-1000, -1000, -1000}, 10000); r != -3000 {
		t.Errorf("expect: %v, actual: %v", -3000, r)
	}
}

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	fmt.Println("sorted", nums)
	
	var sum int
	var mind = 100000
	for i := 0; i < len(nums); i++ {
		left, right := i + 1, len(nums) - 1
		for left < right {
			c := nums[i] + nums[left] + nums[right]
			diff := int(math.Abs(float64(c - target)))
			fmt.Println(diff, mind)
			if diff < mind {
				sum = c
				mind = diff
			}
			if c < target {
				left++
			} else {
				right--
			}
		}
	}
	return sum
}
