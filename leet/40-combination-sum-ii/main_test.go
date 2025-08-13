package main_test

import (
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expect     [][]int
	}{
		{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, expect: [][]int{
			{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6},
		}},
		{candidates: []int{2, 5, 2, 1, 2}, target: 5, expect: [][]int{
			{1, 2, 2}, {5},
		}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := combinationSum2(tt.candidates, tt.target)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	can := make(map[int]int, len(candidates))
	for _, c := range candidates {
		can[c] += 1
	}
	r := make(map[string][]int)
	comb(can, target, []int{}, r)
	arr := make([][]int, 0, len(r))
	for _, v := range r {
		arr = append(arr, v)
	}
	return arr
}

func comb(can map[int]int, target int, current []int, r map[string][]int) {
	if target == 0 {
		arr := make([]int, len(current))
		copy(arr, current)
		slices.Sort(arr)
		r[str(arr)] = arr
		return
	}
	if target < 0 {
		return
	}
	for k, v := range can {
		if v > 0 {
			can[k] -= 1
			current = append(current, k)
			comb(can, target-k, current, r)
			can[k] += 1
			current = current[:len(current)-1]
		}
	}
}

func str(nums []int) string {
	var sb strings.Builder
	for _, n := range nums {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}
