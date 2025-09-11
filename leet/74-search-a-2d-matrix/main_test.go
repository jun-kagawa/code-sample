package main_test

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			expect: true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			expect: false,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 16,
			expect: true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 60,
			expect: true,
		},
		{
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 62, 63, 64}},
			target: 30,
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := searchMatrix(tt.matrix, tt.target); r != tt.expect {
				t.Errorf("expect: %v, actual: %v, target: %v", tt.expect, r, tt.target)
			}
		})
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	arr := make([]int, len(matrix))
	for i, row := range matrix {
		arr[i] = row[0]
	}
	index, ok := search(arr, target)
	if ok {
		return ok
	}
	if _, ok := search(matrix[index], target); ok {
		return true
	}
	return false
}

func search(arr []int, target int) (int, bool) {
	start, end := 0, len(arr)-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] == target {
			return mid, true
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if mid > 0 && arr[mid] > target {
		return mid - 1, false
	}
	return mid, false
}
