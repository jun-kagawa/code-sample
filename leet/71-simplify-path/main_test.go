package main_test

import (
	"strings"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path   string
		expect string
	}{
		{path: "/home/", expect: "/home"},
		{path: "/home", expect: "/home"},
		{path: "/home//foo/", expect: "/home/foo"},
		{path: "/home/user/Documents/../Pictures", expect: "/home/user/Pictures"},
		{path: "/../", expect: "/"},
		{path: "/..../", expect: "/...."},
		{path: "/....", expect: "/...."},
		{path: "/.../a/../b/c/../d/./", expect: "/.../b/d"},
		{path: "/a//b////c/d//././/..", expect: "/a/b/c"},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := simplifyPath(tt.path); r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}

func simplifyPath(path string) string {
	l := len(path)
	r := []string{}
	j := 0
	for i := range l {
		var s string
		if path[i] == '/' && j < i {
			s = path[j+1 : i]
			j = i
		}
		if i == l-1 && path[i] != '/' {
			s = path[j+1:]
		}
		if s == ".." {
			if len(r) > 0 {
				r = r[:len(r)-1]
			}
		} else if s != "" && s != "/" && s != "." {
			r = append(r, s)
		}
	}
	var sb strings.Builder
	if len(r) == 0 {
		sb.WriteString("/")
	} else {
		for _, rr := range r {
			sb.WriteString("/")
			sb.WriteString(rr)
		}
	}
	return sb.String()
}
