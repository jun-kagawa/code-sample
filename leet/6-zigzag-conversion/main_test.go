package main_test

import (
	"strings"
	"testing"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var sb strings.Builder
	for i := range numRows {
		n := (numRows-i)*2 - 2
		for j := i; j < len(s); j += numRows*2 - 2 {
			sb.WriteByte(s[j])
			k := j + n
			if i != 0 && i != numRows-1 && k < len(s) {
				sb.WriteByte(s[k])
			}
		}
	}
	return sb.String()
}

func TestConvert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		expect  string
	}{
		{s: "PAYPALISHIRING", numRows: 3, expect: "PAHNAPLSIIGYIR"},
		{s: "PAYPALISHIRING", numRows: 4, expect: "PINALSIGYAHRPI"},
		{s: "A", numRows: 1, expect: "A"},
		{s: "AA", numRows: 1, expect: "AA"},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := convert(tt.s, tt.numRows)
			if r != tt.expect {
				t.Errorf("failed. expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}
