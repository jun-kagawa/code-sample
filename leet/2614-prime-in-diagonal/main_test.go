package main_test

import (
	"testing"
)

func TestDiagonalPrime(t *testing.T) {
	tests := []struct {
		nums   [][]int
		expect int
	}{
		{
			nums:   [][]int{{1, 2, 3}, {5, 6, 7}, {9, 10, 11}},
			expect: 11,
		},
		{
			nums:   [][]int{{1, 2, 3}, {5, 17, 7}, {9, 11, 10}},
			expect: 17,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := diagonalPrime(tt.nums); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func diagonalPrime(nums [][]int) int {
	r := 0
	j := len(nums) - 1
	for i := range len(nums) {
		if isPrime(nums[i][i]) {
			r = max(r, nums[i][i])
		}
		if isPrime(nums[i][j]) {
			r = max(r, nums[i][j])
		}
		j--
	}
	return r
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
