package main_test

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	if r := strStr("sadbutsad", "sad"); r != 0 {
		t.Errorf("expect: %v, actual: %v", 0, r)
	}
	if r := strStr("leetcode", "leeto"); r != -1 {
		t.Errorf("expect: %v, actual: %v", -1, r)
	}
	if r := strStr("butsad", "sad"); r != 3 {
		t.Errorf("expect: %v, actual: %v", 3, r)
	}
	if r := strStr("butsad", "sadsad"); r != -1 {
		t.Errorf("expect: %v, actual: %v", -1, r)
	}
}

func strStr(heystack string, needle string) int {
	for i := 0; i < len(heystack); i++ {
		if heystack[i] != needle[0] {
			continue;
		}
		limit := i + len(needle)
		if limit > len(heystack) {
			break
		}
		if heystack[i:limit] == needle {
			return i
		}
	}
	return -1
}

