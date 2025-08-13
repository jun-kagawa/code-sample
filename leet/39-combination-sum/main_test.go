package main_test

import (
	"fmt"
	"testing"
	"slices"
	"strings"
	"strconv"

	"github.com/google/go-cmp/cmp"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		expect     [][]int
	}{
		{candidates: []int{2, 3, 6, 7}, target: 7, expect: [][]int{
			{2, 2, 3}, {7},}},
		{candidates: []int{2, 3, 5}, target: 8, expect: [][]int{
			{2, 2, 2, 2}, {2, 3, 3},{3, 5}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test"), func(t *testing.T) {
			r := combinationSum(tt.candidates, tt.target)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}

func comb(arr []int, target int, m *map[string][]int, current []int) {
	if target == 0 {
		slices.Sort(current)
		(*m)[str(current)] = current
		return
	}
	for _, n := range arr {
		if n <= target {
			cur := []int{}
			cur = append(cur, current...)
			cur = append(cur, n)
			comb(arr, target - n, m, cur)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	m := make(map[string][]int)
	comb(candidates, target, &m, []int{})
	r := make([][]int, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
func str(nums []int) string {
	var sb strings.Builder
	for _, n := range nums {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}
