package main_test

import (
	"sort"
	"testing"
)

func TestMinimizeTheDifference(t *testing.T) {
	tests := []struct {
		mat    [][]int
		target int
		expect int
	}{
		{
			mat:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			target: 30,
			expect: 12,
		},
		{
			mat:    [][]int{{1}, {2}, {3}},
			target: 100,
			expect: 94,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := minimizeTheDifference(tt.mat, tt.target); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", r, tt.expect)
			}
		})
	}
}

func minimizeTheDifference(mat [][]int, target int) int {
	base := 0
	magic := make([][]int, len(mat))
	for i, row := range mat {
		sort.Ints(row)
		base += row[0]
		for j := 1; j < len(row); j++ {
			if row[j] != row[j-1] {
				magic[i] = append(magic[i], row[j]-row[0])
			}
		}
	}
	minDiff := abs(target - base)
	if base >= target {
		return minDiff
	}
	hash := make(map[int]bool, 871)
	hash[base] = true

	for _, row := range magic {
		more := make([]int, 0)
		for _, cell := range row {
			for basic := range hash {
				v := basic + cell
				if v > 870 {
					continue
				}
				if hash[v] == false {
					more = append(more, v)
				}
				if v == target {
					return 0
				}
			}
		}
		for _, mv := range more {
			hash[mv] = true
		}
	}
	for k := range hash {
		if cur := abs(target - k); cur < minDiff {
			minDiff = cur
		}
	}
	return minDiff
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return n * -1
}
