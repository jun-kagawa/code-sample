package main_test

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expect: 49},
		{input: []int{1, 1}, expect: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input: %v", tt.input), func(t *testing.T) {
			if r := maxArea(tt.input); r != tt.expect {
				t.Errorf("expect %v, actual %v", tt.expect, r)
			}
		})
	}

}

func maxArea(height []int) int {
	m := 0
	i := 0
	j := len(height) - 1
	for i < j {
		y := min(height[i], height[j])
		x := j - i
		area := y * x
		if area > m {
			m = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}
