package main_test

import (
	// "math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		n      int
		expect []int
	}{
		{
			n:      2,
			expect: []int{0, 1, 3, 2},
		},
		{
			n:      1,
			expect: []int{0, 1},
		},
		{
			n:      3,
			expect: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := grayCode(tt.n)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

/**
----
00
01
11
10
----
000
000
000
---
001
000
001
---
010
001
011
---
011
001
010
---
100
010
110
---
101
010
111
---
110
011
101
---
111
011
100
----
**/

func grayCode(n int) []int {
	r := make([]int, 0, 1<<n)
	for i := range 1 << n {
		r = append(r, i^(i>>1))
	}
	return r
}
