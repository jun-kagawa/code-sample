package main_test

import (
	//	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1     string
		s2     string
		s3     string
		expect bool
	}{
		{
			s1:     "aabcc",
			s2:     "dbbca",
			s3:     "aadbbcbcac",
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := isInterleave(tt.s1, tt.s2, tt.s3); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n]
}
