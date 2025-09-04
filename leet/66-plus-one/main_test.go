package main_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		expect []int
	}{
		{digits: []int{1, 2, 3}, expect: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, expect: []int{4, 3, 2, 2}},
		{digits: []int{9}, expect: []int{1, 0}},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := plusOne(tt.digits)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("diff: %v, actual: %v, expect: %v", diff, r, tt.expect)
			}
		})
	}
}

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			return digits
		} else {
			digits[i] = 0
		}
	}
	r := make([]int, len(digits)+1)
	r[0] = 1
	copy(r[1:], digits)
	return r
}
