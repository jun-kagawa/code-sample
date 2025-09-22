package main_test

import (
	"slices"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		nums   []int
		expect [][]int
	}{
		{
			nums:   []int{1, 2, 2},
			expect: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			nums:   []int{0},
			expect: [][]int{{}, {0}},
		},
		{
			nums: []int{4, 4, 4, 1, 4},
			expect: [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4},
				{4, 4}, {4, 4, 4}, {4, 4, 4, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := subsetsWithDup2(tt.nums)

			sort.Slice(r, func(i, j int) bool {
				a, b := r[i], r[j]
				// 要素を順に比較
				for k := 0; k < len(a) && k < len(b); k++ {
					if a[k] < b[k] {
						return true
					}
					if a[k] > b[k] {
						return false
					}
				}
				// ここまで一致した場合は「長さの短い方が小さい」
				return len(a) < len(b)
			})
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v, diff: %v", tt.expect, r, diff)
			}
		})
	}
}

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	r := make(map[string][]int)
	back(nums, []int{}, &r, len(nums))

	rr := [][]int{}
	for _, arr := range r {
		fmt.Println(arr)
		rr = append(rr, arr)
	}

	return rr
}

func back(nums, cur []int, r *map[string][]int, l int) {
	c := make([]int, len(cur))
	copy(c, cur)
	var sb strings.Builder
	for _, n := range c {
		sb.WriteString(strconv.Itoa(n))
	}
	(*r)[sb.String()] = c
	if len(nums) == 0 || len(cur) == l {
		return
	}
	for i, n := range nums {
		cur = append(cur, n)
		nn := make([]int, len(nums))
		copy(nn, nums)
		nn = nn[i+1:]
		back(nn, cur, r, l)
		cur = cur[:len(cur)-1]
	}
}


func subsetsWithDup2(nums []int) [][]int {
	r := [][]int{}
	cnt := make(map[int]int)
	for _, n := range nums {
		cnt[n]++
	}
	back2(&r, cnt, []int{})
	return r
}

func back2(r *[][]int, cnt map[int]int, cur []int) {
	c := make([]int, len(cur))
	copy(c, cur)
	(*r) = append(*r, c)

	for k := range cnt {
		if cnt[k] == 0 || (len(cur) > 0 && cur[len(cur)-1] > k){
			continue
		}
		cnt[k]--
		cur = append(cur, k)

		back2(r, cnt, cur)

		cnt[k]++
		cur = cur[:len(cur)-1]
	}
}
