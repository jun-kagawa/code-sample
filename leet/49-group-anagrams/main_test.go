package main_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var opt cmp.Option = cmpopts.SortSlices(func(i, j int) bool {
	return i < j
})

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs   []string
		expect [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expect: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			strs:   []string{""},
			expect: [][]string{{""}},
		},
		{
			strs:   []string{"a"},
			expect: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := groupAnagrams(tt.strs)
			fmt.Println(r)
			diff := cmp.Diff(r, tt.expect, opt)
			if diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}
func groupAnagrams2(strs []string) [][]string {
	m := make(map[string][]int)
	for i, str := range strs {
		ru := []rune(str)
		slices.Sort(ru)
		s := string(ru)
		m[s] = append(m[s], i)
	}
	r := make([][]string, 0, len(strs))
	for _, v := range m {
		arr := make([]string, 0, len(v))
		for _, idx := range v {
			arr = append(arr, strs[idx])
		}
		r = append(r, arr)
	}
	return r
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[int][]string)
	for i, str := range strs {
		ru := []rune(str)
		var bit int
		for _, rr := range ru {
			bit += (1 << (rr - 'a'))
		}
		m[bit] = append(m[bit], strs[i])
	}
	r := make([][]string, 0, len(strs))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for i, str := range strs {
		ru := []rune(str)
		slices.Sort(ru)
		s := string(ru)
		m[s] = append(m[s], strs[i])
	}
	r := make([][]string, 0, len(strs))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
