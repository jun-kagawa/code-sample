package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		expect []int
	}{
		{
			a:      []int{1, 3, 2, 4},
			b:      []int{3, 1, 2, 4},
			expect: []int{0, 2, 3, 4},
		},
		{
			a:      []int{2, 3, 1},
			b:      []int{3, 1, 2},
			expect: []int{0, 1, 3},
		},
		{
			a:      []int{1, 2, 3},
			b:      []int{1, 2, 3},
			expect: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := findThePrefixCommonArray(tt.a, tt.b)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func findThePrefixCommonArray(A []int, B []int) []int {
	l := len(A)
	res := make([]int, l + 1)
	tmp := make([]int, l)
	s := 0

	for i := range l {
		res[A[i]]++
		if res[A[i]] == 2 {
			s++
		}

		res[B[i]]++
		if res[B[i]] == 2 {
			s++
		}
		tmp[i] = s
	}
	return tmp
}

func findThePrefixCommonArray1(A []int, B []int) []int {
	r := make([]int, len(A))
	a := make(map[int]struct{})
	b := make(map[int]struct{})
	for i := range r {
		a[A[i]] = struct{}{}
		b[B[i]] = struct{}{}
		c := 0
		for k := range a {
			if _, ok := b[k]; ok {
				c++
			}
		}
		r[i] = c
	}
	return r
}
