package main_test

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		expect   int
	}{
		{dividend: 10, divisor: 3, expect: 3},
		{dividend: 7, divisor: -3, expect: -2},
		{dividend: -7, divisor: -3, expect: 2},
		{dividend: -7, divisor: 3, expect: -2},
		{dividend: 0, divisor: 1, expect: 0},
		{dividend: -2147483648, divisor: -1, expect: 2147483647},
		{dividend: -2147483648, divisor: 4, expect: -536870912},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := divide(tt.dividend, tt.divisor); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

var upper = int(math.Pow(2, 31) - 1)
var under = int(math.Pow(-2, 31))

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	negative := (dividend < 0) != (divisor < 0)
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	c := 0
	for dividend > divisor {
		sub := divisor
		add := 1
		for dividend >= sub << 1 {
			sub <<= 1
			add <<= 1
		}
		c += add
		dividend -= sub
	}
	if negative {
		return -c
	}
	return c
}

func divide1(dividend int, divisor int) int {
	mDividend := false
	if dividend < 0 {
		dividend = int(math.Abs(float64(dividend)))
		mDividend = true
	}
	mDivisor := false
	if divisor < 0 {
		divisor = int(math.Abs(float64(divisor)))
		mDivisor = true
	}
	c := 0
	if divisor == 1 {
		c = dividend
		dividend = 0
	}
	for {
		dividend -= divisor
		if dividend < 0 {
			break
		}
		c++
	}
	var r int
	if (mDivisor && mDividend) || (!mDivisor && !mDividend) {
		r = c
	} else {
		r = 0 - c
	}
	if r > upper {
		return upper
	}
	if r < under {
		return under
	}
	return r
}

