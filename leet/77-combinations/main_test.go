package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		expect [][]int
	}{
		{n: 4, k: 2, expect: [][]int{
			{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
		}},
		{n: 1, k: 1, expect: [][]int{{1}}},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := combine(tt.n, tt.k)
			diff := cmp.Diff(tt.expect, r)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func combine2(n int, k int) [][]int {
	var counter int
	a, b := 1, 1
	for i := range k {
		a *= (n - i)
	}
	for i := 1; i <= k; i++ {
		b *= i
	}
	l := a / b
	r := make([][]int, 0, l)
	for range l {
		r = append(r, make([]int, k))
	}

	arr := make([]int, 0, k)
	for i := range k {
		arr = append(arr, i+1)
	}

	for {
		copy(r[counter], arr)
		counter++

		i := k - 1
		for i >= 0 && arr[i] == n-k+i+1 {
			i--
		}
		if i < 0 {
			break
		}
		arr[i]++
		for j := i + 1; j < k; j++ {
			arr[j] = arr[j-1] + 1
		}
	}
	return r
}

func combine(n int, k int) [][]int {
	a, b := 1, 1
	for i := range k {
		a *= (n - i)
	}
	for i := 1; i <= k; i++ {
		b *= i
	}
	l := a / b
	r := make([][]int, 0, l)
	arr := make([]int, 0, n)
	for i := range n {
		arr = append(arr, i+1)
	}

	back(arr, k, []int{}, &r)
	return r
}

func back(arr []int, k int, cur []int, r *[][]int) {
	if len(cur) == k {
		(*r) = append(*r, cur)
		return
	}
	if len(arr) == 0 {
		return
	}
	for i := range arr {
		cur = append(cur, arr[i])
		c := make([]int, len(cur))
		copy(c, cur)
		back(arr[i+1:], k, c, r)
		cur = cur[:len(cur)-1]
	}
}
