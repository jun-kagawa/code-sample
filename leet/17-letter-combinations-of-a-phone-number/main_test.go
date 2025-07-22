package main_test

import (
	"strconv"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	r := letterCombinations("23")
	if len(r) != 9 {
		t.Errorf("unexpected length. actual: %v, expect: %v", len(r), 9)
	}
	expect := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	for i, e := range expect {
		if e != r[i] {
			t.Errorf("expect: %v, actual: %v", e, r[i])
		}
	}
	r = letterCombinations("234")
	if len(r) != 27 {
		t.Errorf("unexpected length. actual: %v, expect: %v", len(r), 9)
	}
	expect = []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	for i, e := range expect {
		if e != r[i] {
			t.Errorf("expect: %v, actual: %v", e, r[i])
		}
	}

}

func letterCombinations(digits string) []string {
	m := make(map[string][]string)
	ss := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	for i, s := range ss {
		m[strconv.Itoa(i+2)] = s
	}

	strs := make([][]string, 0, len(digits))
	for _, d := range digits {
		if ss, ok := m[string(d)]; ok {
			strs = append(strs, ss)
		}
	}
	r := []string{}
	if len(strs) == 0 {
		return r
	}
	for _, s := range strs[0] {
		r = append(r, s)
	}
	for _, ss := range strs[1:] {
		rrr := []string{}
		for _, rr := range r {
			for _, s := range ss {
				rrr = append(rrr, rr+s)
			}
		}
		r = rrr
	}
	return r
}

