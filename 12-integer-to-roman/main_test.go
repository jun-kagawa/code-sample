package main_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	if r := intToRoman(3749); r != "MMMDCCXLIX" {
		t.Errorf("actual: %v, expect: %v", r, "MMMDCCXLIX")
	}
	if r := intToRoman(58); r != "LVIII" {
		t.Errorf("actual: %v, expect: %v", r, "LVIII")
	}
	if r := intToRoman(1994); r != "MCMXCIV" {
		t.Errorf("actual: %v, expect: %v", r, "MCMXCIV")
	}
	if r := intToRoman(10); r != "X" {
		t.Errorf("actual: %v, expect: %v", r, "X")
	}
}

func genString(c int, sb *strings.Builder, high, mid, low string) {
	if c > 0 {
		if c == 9 {
			sb.WriteString(low)
			sb.WriteString(high)
		} else if c >= 5 {
			sb.WriteString(mid)
			for range c - 5 {
				sb.WriteString(low)
			}
		} else if c == 4 {
			sb.WriteString(low)
			sb.WriteString(mid)
		} else {
			for range c {
				sb.WriteString(low)
			}
		}
	}
}

func intToRoman(num int) string {
	a := num / 1000
	b := num % 1000 / 100
	c := num % 1000 % 100 / 10
	d := num % 1000 % 100 % 10 / 1
	fmt.Println(a, b, c, d)
	var sb strings.Builder
	if a > 0 {
		for range a {
			sb.WriteString("M")
		}
	}
	genString(b, &sb, "M", "D", "C")
	genString(c, &sb, "C", "L", "X")
	genString(d, &sb, "X", "V", "I")
	return sb.String()
}
