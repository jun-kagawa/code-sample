package main_test

import (
	"testing"
)

func TestNumberOfStableArrays(t *testing.T) {
	tests := []struct{
		one int
		zero int
		limit int
		expect int
	}{
		{
			one: 1,
			zero: 1,
			limit: 2,
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run("test", func (t *testing.T) {
			if r := numberOfStableArrays(tt.zero, tt.one, tt.limit); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

const mod = 1_000_000_007

var dp [202][202][2]int

func solve(one, zero int, ms bool, limit int) int {
	if one == 0 && zero == 0 {
		return 1
	}

	var msIdx int
	if ms {
		msIdx = 1
	} else {
		msIdx = 0
	}

	if dp[one][zero][msIdx] != -1 {
		return dp[one][zero][msIdx]
	}

	ans := 0
	if ms {
		for i := 1; i <= min(one, limit); i++ {
			ans = (ans + solve(one-i, zero, false, limit)) % mod
		}
	} else {
		for i := 1; i <= min(zero, limit); i++ {
			ans = (ans + solve(one, zero-i, true, limit)) % mod
		}
	}

	dp[one][zero][msIdx] = ans % mod
	return dp[one][zero][msIdx]
}

func numberOfStableArrays(zero, one, limit int) int {
	// dp初期化
	for i := 0; i <= 201; i++ {
		for j := 0; j <= 201; j++ {
			dp[i][j][0] = -1
			dp[i][j][1] = -1
		}
	}

	a := solve(one, zero, true, limit) % mod
	b := solve(one, zero, false, limit) % mod
	return (a + b) % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

