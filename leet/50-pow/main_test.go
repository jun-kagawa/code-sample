package main_test

import (
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct{
		x float64
		n int
		expect float64
	}{
		{x: 2.000, n: 10, expect: 1024.0000},
		{x: 2.100, n: 3, expect: 9.26100},
		{x: 2.000, n: -2, expect: 0.25000},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := myPow(tt.x, tt.n); r != tt.expect {
				t.Errorf("actual: %v, expect: %v", r, tt.expect)
			}
		})
	}
}

func myPow(x float64, n int) float64 {
	nn := int64(n)
	if nn < 0 {
		x = 1 / x
		nn = -nn
	}
	r := int64(helper(x, nn) * 10000)
	return float64(r) / 10000
}

func helper(x float64, n int64) float64 {
	if n == 0 {
		return 1
	}
	half := helper(x, n / 2)
	r := half * half
	if n % 2 == 1 {
		r *= x
	}
	return r
}
