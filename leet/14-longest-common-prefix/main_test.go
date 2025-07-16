package main_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	if r := longestCommonPrefix([]string{"flower","flow","flight"}); r != "fl" {
		t.Errorf("expect: %v, actual: %v", "fl", r)
	}
	if r := longestCommonPrefix([]string{"dog","racecar","car"}); r != "" {
		t.Errorf("expect: %v, actual: %v", "", r)
	}
	if r := longestCommonPrefix([]string{"dog"}); r != "dog" {
		t.Errorf("expect: %v, actual: %v", "dog", r)
	}
	if r := longestCommonPrefix([]string{}); r != "" {
		t.Errorf("expect: %v, actual: %v", "", r)
	}
	if r := longestCommonPrefix([]string{"ab", "a"}); r != "a" {
		t.Errorf("expect: %v, actual: %v", "a", r)
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	p := strs[0]
	for i := 1; i < len(strs); i++ {
		if p == "" {
			return p
		}
		var sb strings.Builder
		s := min(len(p), len(strs[i]))
		fmt.Println("s", s)
		for j := 0; j < s; j++ {
			if p[j] == strs[i][j] {
				sb.WriteByte(p[j])
			} else {
				break
			}
		}
		p = sb.String()
	}
	return p
}
