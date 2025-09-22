package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{
			s:      "12",
			expect: 2,
		},
		{
			s:      "226",
			expect: 3,
		},
		{
			s:      "06",
			expect: 0,
		},
		{
			s:      "111111111111111111111111111111111111111111111",
			expect: 1836311903,
		},
		{
			s:      "11111",
			expect: 8,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := numDecodings(tt.s); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
				fmt.Println(r)
			}
		})
	}
}

func numDecodings(s string) int {
	l := len(s)
	r := make([]int, l)
	for i := range l {
		d := s[i] - '0'
		if 0 < d && d < 10 {
			if i == 0 {
				r[i] = 1
			} else {
				r[i] += r[i-1]
			}
		}
		if i == 0 {
			continue
		}
		dd, _ := strconv.Atoi(s[i-1 : i+1])
		if 9 < dd && dd < 27 {
			if i == 1 {
				r[i] += 1
			} else {
				r[i] += r[i-2]
			}
		}
	}
	return r[l-1]
}
