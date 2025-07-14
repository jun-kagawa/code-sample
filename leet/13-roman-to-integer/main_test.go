package main_test

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		i string
		e int
	}{
		{i: "III", e: 3},
		{i: "LVIII", e: 58},
		{i: "MCMXCIV", e: 1994},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input: %v", tt.i), func(t *testing.T) {
			if r := romanToInt(tt.i); r != tt.e {
				t.Errorf("expect: %v, actual: %v", tt.e, r)
			}
		})
	}
}

func romanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		var next string
		if i + 1 < len(s) {
			next = string(s[i+1])
		}
		ss := string(s[i])
		if ss == "M" {
			sum += 1000
		} else if ss == "D" {
			sum += 500
		} else if ss == "C" {
			sum += 100
			if next == "D" {
				sum += 300
				i++
			} else if next == "M" {
				sum += 800 
				i++
			}
		} else if ss == "L" {
			sum += 50
		} else if ss == "X" {
			sum += 10
			if next == "L" {
				sum += 30
				i++
			} else if next == "C" {
				sum += 80
				i++
			}
		} else if ss == "V" {
			sum += 5
		} else if ss == "I" {
			sum += 1
			if next == "V" {
				sum += 3
				i++
			} else if next == "X" {
				sum += 8
				i++
			}
		}
	}
	return sum
}
