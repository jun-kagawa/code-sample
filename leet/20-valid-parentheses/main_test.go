package main_test

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{input: "()", expect: true},
		{input: "()[]{}", expect: true},
		{input: "(]", expect: false},
		{input: "([])", expect: true},
		{input: "([)]", expect: false},
		{input: "]", expect: false},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := isValid(tt.input)
			if r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func isValid(s string) bool {
	o := map[string]bool{
		"(": true,
		"{": true,
		"[": true,
	}
	c := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		ss := string(s[i])
		if _, ok := o[ss]; ok {
			stack = append(stack, ss)
		} else if v, ok := c[ss]; ok {
			if len(stack) < 1 {
				return false
			}
			pop := stack[len(stack)-1]
			if pop != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
