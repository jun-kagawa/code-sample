package main_test

import (
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	return v2(s)
}

func v1(s string) int {
	var c int
	var ma int
	m := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				m = make(map[byte]struct{})
				ma = max(c, ma)
				c = 0
				break
			}
			m[s[j]] = struct{}{}
			c++
		}
	}
	return max(c, ma)
}
func v2(s string) int {
	var lastPos [256]int
	left := 0
	maxLen := 0
	for right := 0; right < len(s); right++ {
		char := s[right]
		
		if lastPos[char] > left {
			left = lastPos[char]
		}

		lastPos[char] = right + 1
		
		maxLen = max(maxLen, right - left + 1)
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s string
		i int
	}{
		{s: "abcabcbb", i: 3},
		{s: "bbbb", i: 1},
		{s: "pwwkew", i: 3},
		{s: " ", i: 1},
		{s: "dvdf", i: 3},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := lengthOfLongestSubstring(tt.s)
			if r != tt.i {
				t.Errorf("expect: %d, actual: %d", tt.i, r)
			}
		})
	}
}
