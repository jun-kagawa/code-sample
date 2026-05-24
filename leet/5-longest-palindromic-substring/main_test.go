package main_test

import (
	"testing"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}
		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{
			s:      "babad",
			expect: "bab", // "aba" も正解ですが、テストの期待値に合わせています
		},
		{
			s:      "cbbd",
			expect: "bb",
		},
		{
			s:      "a",
			expect: "a",
		},
		{
			s:      "ac",
			expect: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			r := longestPalindrome(tt.s)
			if r != tt.expect {
				t.Errorf("failed for %s. expect: %s, actual: %s", tt.s, tt.expect, r)
			}
		})
	}
}
