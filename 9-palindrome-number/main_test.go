package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		num    int
		expect bool
	}{
		{num: 121, expect: true},
		{num: -121, expect: false},
		{num: 10, expect: false},
		{num: 0, expect: true},
		{num: 1111, expect: true},
		{num: 11111, expect: true},
		{num: 13131, expect: true},
		{num: -13131, expect: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("input: %v", tt.num), func(t *testing.T) {
			if r := isPalindrome(tt.num); r != tt.expect {
				t.Errorf("actual %v, expected %v", r, tt.expect)
			}
		})
	}
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1
	for left < right {
		fmt.Println("left", s[left], "right", s[right])
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
