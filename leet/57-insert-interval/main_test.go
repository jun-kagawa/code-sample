package main_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		expect      [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expect:      [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expect:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expect:      [][]int{{5, 7}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			expect:      [][]int{{1, 7}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := insert(tt.intervals, tt.newInterval)
			fmt.Println(r)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	if l == 0 {
		return [][]int{newInterval}
	}
	for i, val := range intervals {
		if val[0] > newInterval[0] {
			t := make([][]int, len(intervals[i:]))
			copy(t, intervals[i:])
			intervals = append(intervals[:i], newInterval)
			intervals = append(intervals, t...)
			break
		}
		if i == l-1 {
			intervals = append(intervals, newInterval)
		}
	}
	r := [][]int{}
	t := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if t[1] >= intervals[i][0] {
			t[1] = max(t[1], intervals[i][1])
		} else {
			r = append(r, t)
			t = intervals[i]
		}
	}
	r = append(r, t)
	return r
}
