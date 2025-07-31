package main_test

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	arr := []int{3, 2, 2, 3}
	if r := removeElement(arr, 3); r != 2 {
		t.Errorf("expect: %v, actual: %v", 2, r)
	}
	fmt.Println(arr)
	arr = []int{0, 1, 2, 2, 3, 0, 4, 2}
	if r := removeElement(arr, 2); r != 5 {
		t.Errorf("expect: %v, actual: %v", 5, r)
	}
	fmt.Println(arr)
}

func removeElement(nums []int, val int) int {
	var c int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			c++
			continue
		}
		for j := len(nums) - 1; i < j; j-- {
			if nums[j] != val {
				nums[i], nums[j] = nums[j], nums[i]
				c++
				break
			}
		}
	}
	return c
}
