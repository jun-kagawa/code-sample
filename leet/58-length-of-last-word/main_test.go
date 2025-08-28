package main_test

import (
	"testing"
)

func TestLenghOfLastWord(t *testing.T) {
	tests := []struct{
		s string
		expect int
	}{
		{s: "Hello World", expect: 5},
		{s: "   fly me   to   the moon  ", expect: 4},
		{s: "luffy is still joyboy", expect: 6},
		{s: "a", expect: 1},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := lengthOfLastWord(tt.s); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func lengthOfLastWord(s string) int {
	var last string
	var j int
	for i := range len(s) {
		if s[i] == ' ' {
			ss := s[j:i]
			j = i + 1
			if ss != "" {
				last = ss
			}
		}
	}
	if j <= len(s) - 1 {
		last = s[j:]
	}
	return len(last)
}
