package main_test

import (
	"strings"
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		expect string
	}{
		{a: "11", b: "1", expect: "100"},
		{a: "1", b: "1", expect: "10"},
		{a: "0", b: "1", expect: "1"},
		{a: "0", b: "0", expect: "0"},
		{a: "1010", b: "1011", expect: "10101"},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := addBinary(tt.a, tt.b); r != tt.expect {
				t.Errorf("expect: %v, actual:%v", tt.expect, r)
			}
		})
	}
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	diff := max(la, lb) - min(la, lb)
	var sb strings.Builder
	for range diff {
		sb.WriteRune('0')
	}
	if la > lb {
		sb.WriteString(b)
		b = sb.String()
	} else if lb > la {
		sb.WriteString(a)
		a = sb.String()
	}
	var r strings.Builder
	l := len(a)
	t := '0'
	for i := l - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if t == '1' {
				r.WriteRune('1')
			} else {
				r.WriteRune('0')
			}
			t = '1'
		} else if a[i] == '0' && b[i] == '0' {
			if t == '1' {
				r.WriteRune('1')
			} else {
				r.WriteRune('0')
			}
			t = '0'
		} else {
			if t == '1' {
				r.WriteRune('0')
				t = '1'
			} else {
				r.WriteRune('1')
				t = '0'
			}
		}
	}
	r.WriteRune(t)
	return reverse(r.String())
}

func reverse(s string) string {
	var r strings.Builder
	l := len(s)
	f := true
	for i := l - 1; i >= 0; i-- {
		if s[i] == '0' && f {
			continue
		}
		if s[i] == '1' && f {
			f = false
		}
		r.WriteByte(s[i])
	}
	if r.String() == "" {
		return "0"
	}
	return r.String()
}
