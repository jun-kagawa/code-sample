package main_test

import (
	"testing"
)

func TestMaximumEnergy(t *testing.T) {
	tests := []struct {
		energy []int
		k      int
		expect int
	}{
		{energy: []int{5, 2, -10, -5, 1}, k: 3, expect: 3},
		{energy: []int{-2, -3, -1}, k: 2, expect: -1},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := maximumEnergy(tt.energy, tt.k); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func maximumEnergy(energy []int, k int) int {
	l := len(energy) - 1
	for i := l; i >= 0; i-- {
		j := i + k
		if j > l {
			continue
		}
		energy[i] += energy[j]
	}
	return m(energy)
}

func m(nums []int) int {
	maxx := nums[0]
	for _, n := range nums {
		if n > maxx {
			maxx = n
		}
	}
	return maxx
}
