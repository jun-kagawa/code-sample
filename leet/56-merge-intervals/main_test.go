package main_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expect    [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expect:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expect:    [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 5}},
			expect:    [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 4}, {2, 3}},
			expect:    [][]int{{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := merge(tt.intervals)
			fmt.Println(r)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	r := [][]int{}
	t := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if t[1] >= intervals[i][0] {
			t[1] = max(intervals[i][1], t[1])
		} else {
			r = append(r, t)
			t = intervals[i]
		}
	}
	r = append(r, t)
	return r
}
