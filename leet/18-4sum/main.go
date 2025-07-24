package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	/*
		r := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
		fmt.Println(r)

		r = fourSum([]int{2, 2, 2, 2, 2}, 8)
		fmt.Println(r)

		r = fourSum([]int{2}, 2)
		fmt.Println(r)

		r = fourSum([]int{-3, -1, 0, 2, 4, 5}, 0)
		fmt.Println(r)

		r = fourSum([]int{-3, -1, 0, 2, 4, 5}, 2)
		fmt.Println(r)
		r := fourSum([]int{-1, 0, 1, 2, -1, -4}, -1)
		fmt.Println(r)
	*/
	r := fourSum([]int{-1, 2, 2, -5, 0, -1, 4}, 3)
	fmt.Println(r)

}

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)

	a := 0
	b := a + 1
	c := len(nums) - 2
	d := len(nums) - 1

	m := make(map[string][]int)

	for a+1 < d-1 {
		for b < c {
			sum := nums[a] + nums[b] + nums[c] + nums[d]

			if sum == target {
				s := join([]int{nums[a], nums[b], nums[c], nums[d]})
				m[s] = []int{nums[a], nums[b], nums[c], nums[d]}
				b++
			} else if sum > target {
				c--
			} else {
				b++
			}
		}
		d--
		c = d - 1
		b = a + 1
		if d-1 <= a+1 {
			a++
			b = a + 1
			c = len(nums) - 2
			d = len(nums) - 1
		}
	}
	r := [][]int{}
	for _, mm := range m {
		r = append(r, mm)
	}
	return r
}

func join(nums []int) string {
	ss := make([]string, 0, len(nums))
	for _, n := range nums {
		ss = append(ss, strconv.Itoa(n))
	}
	return strings.Join(ss, ",")
}
