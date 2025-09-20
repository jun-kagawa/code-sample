package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		num1   []int
		m      int
		num2   []int
		n      int
		expect []int
	}{
		{
			num1:   []int{1, 2, 3, 0, 0, 0},
			m:      3,
			num2:   []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
		{
			num1:   []int{1},
			m:      1,
			num2:   []int{},
			n:      0,
			expect: []int{1},
		},
		{
			num1:   []int{0},
			m:      0,
			num2:   []int{1},
			n:      1,
			expect: []int{1},
		},
		{
			num1:   []int{2, 5, 6, 0, 0, 0},
			m:      3,
			num2:   []int{1, 2, 3},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			merge(tt.num1, tt.m, tt.num2, tt.n)
			diff := cmp.Diff(tt.num1, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, tt.num1)
			}
		})
	}
}

func merge(num1 []int, m int, num2 []int, n int) {
	r := make([]int, m + n)
	i, j := 0, 0
	for i < m && j < n {
		if num1[i] > num2[j] {
			r[i+j] = num2[j]
			j++
		} else {
			r[i+j] = num1[i]
			i++
		}
	}
	if i != m {
		copy(r[i+j:], num1[i:])
	} else if j != n {
		copy(r[i+j:], num2[j:])
	}
	copy(num1, r)
}
