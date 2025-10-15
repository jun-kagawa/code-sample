package main_test

import (
	"fmt"
	"slices"
	"testing"
)

func TestMinDeletions(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{s: "aab", expect: 0},
		{s: "aaabbbcc", expect: 2},
		{s: "ceabaacb", expect: 2},
		{s: "bbcebab", expect: 2},
		{s: "abcabc", expect: 3},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := minDeletions(tt.s); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func minDeletions(s string) int {
	h := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		h[s[i]]++
	}
	values := make([]int, 0, len(h))
	for _, v := range h {
		values = append(values, v)
	}
	slices.SortFunc(values, func(a, b int) int {
		return b - a
	})
	count := 0
	fmt.Println(values)
	for i := 1; i < len(values); i++ {
		if values[i-1] > values[i] {
			continue;
		}
		diff := abs(values[i-1] - values[i]) + 1
		if values[i-1] == 0 {
			diff = values[i]
		}
		fmt.Println(diff)
		values[i] -= diff
		count += diff
	}
	return count
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
