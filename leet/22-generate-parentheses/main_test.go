package main_test

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	r := generateParenthesis(4)
	fmt.Println(r)
}

var number int

func generateParenthesis(n int) []string {
	number = n
	comb := []string{}
	gen("", 0, 0, &comb)
	return comb
}

func gen(current string, open, close int, comb *[]string) {
	fmt.Println("current: ", current)
	if len(current) == number * 2 {
		*comb = append(*comb, current)
		return
	}
	if open < number {
		gen(current+"(", open+1, close, comb)
	}
	if close < open {
		gen(current+")", open, close+1, comb)
	}
}

