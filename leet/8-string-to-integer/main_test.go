package main_test

import (
	"strings"
	"testing"
)

var digits [256]bool
var maxValue = (2 << 30) - 1
var minValue = (-2 << 30)

func init() {
	digit := []int{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, d := range digit {
		digits[d] = true
	}
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	var sb strings.Builder
	var m bool
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			m = true
		}
		s = s[1:]
	}
	for _, b := range s {
		if digits[b] {
			sb.WriteRune(b)
		} else {
			break
		}
	}
	str := sb.String()
	if str == "" {
		return 0
	}
	val := 0
	for _, s := range str{
		val = val * 10 + int(s - '0')
        if m && -val <= minValue {
            return minValue
        }
        if !m && val >= maxValue {
            return maxValue
        }
	}
	if m {
		return -val
	}
	return val
}

func TestMyAtoi(t *testing.T) {
	tests := []struct{
		s string
		o int
	}{
		{s: "42", o: 42},
		{s: " -42", o: -42},
		{s: "1137c0d3", o: 1137},
		{s: "0-1", o: 0},
		{s: "1-1", o: 1},
		{s: "words and 987", o: 0},
		{s: "", o: 0},
		{s: "-91283472332", o: -2147483648},
	}

	for _, tt := range tests {
		if r := myAtoi(tt.s); r != tt.o {
			t.Errorf("failed, actual: %v, expect: %v", r, tt.o)
		}
	}
}

