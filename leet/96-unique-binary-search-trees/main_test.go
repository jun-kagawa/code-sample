package main_test

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{
			n: 3, expect: 5,
		},
		{
			n: 1, expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := numTrees(tt.n); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
				fmt.Println(r)
			}
		})
	}

}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	memo := make([][]int, n+2)
	for i := range memo {
		memo[i] = make([]int, n+2)
	}
	return gen(1, n, memo)
}

func gen(start, end int, memo [][]int) int {
	if start > end {
		return 1
	}
	if e := memo[start][end]; e != 0 {
		return e
	}

	var rr int
	for i := start; i <= end; i++ {
		l := gen(start, i-1, memo)
		r := gen(i+1, end, memo)
		rr += r * l
	}
	memo[start][end] = rr
	return rr
}
