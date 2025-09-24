package main_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		s      string
		expect []string
	}{
		{
			s:      "25525511135",
			expect: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			s:      "0000",
			expect: []string{"0.0.0.0"},
		},
		{
			s:      "101023",
			expect: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := restoreIpAddresses(tt.s)
			diff := cmp.Diff(r, tt.expect)
			if diff != "" {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
			fmt.Println(diff)
		})
	}
}

func restoreIpAddresses(s string) []string {
	r := []string{}
	back(s, []string{}, &r)
	return r
}

func back(s string, cur []string, r *[]string) {
	if len(cur) == 4 && len(s) == 0 {
		ss := strings.Join(cur, ".")
		(*r) = append(*r, ss)
		return
	}
	for i := range 3 {
		if len(s) > i {
			ss := s[:i+1]
			if len(ss) > 1 && ss[0] == '0' {
				continue
			}
			if n, _ := strconv.Atoi(ss); n < 256 {
				cur = append(cur, ss)
				back(s[i+1:], cur, r)
				cur = cur[:len(cur)-1]
			}
		}
	}
}
