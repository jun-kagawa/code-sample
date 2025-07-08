package main_test

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	if r := myAtoi("42"); r != 42 {
		t.Errorf("actual: %v, expected %v", r, 42)
	}
	if r := myAtoi("   -042"); r != -42 {
		t.Errorf("actual: %v, expected %v", r, 42)
	}
	if r := myAtoi("1337c0d3"); r != 1337 {
		t.Errorf("actual: %v, expected %v", r, 1337)
	}
	if r := myAtoi("0-1"); r != 0 {
		t.Errorf("actual: %v, expected %v", r, 0)
	}
	if r := myAtoi("words and 987"); r != 0 {
		t.Errorf("actual: %v, expected %v", r, 0)
	}
	if r := myAtoi("-91283472332"); r != -2147483648 {
		t.Errorf("actual: %v, expected %v", r, -2147483648)
	}
	if r := myAtoi("-+12"); r != 0 {
		t.Errorf("actual: %v, expected %v", r, 0)
	}
	if r := myAtoi("20000000000000000000"); r != 2147483647 {
		t.Errorf("actual: %v, expected %v", r, 2147483647)
	}

}

var m = make(map[string]bool)
var digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "+"}
var minx = int(math.Pow(-2, 31))
var maxx = int(math.Pow(2, 31) - 1)

func myAtoi(s string) int {
	for _, i := range digits {
		m[i] = true
	}
	var isNegative bool
	var isDigit bool
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		ss := string(s[i])
		_, ok := m[ss]
		if !ok && !isDigit {
			if ss != " " {
				b.WriteString("0")
				break
			}
			continue
		}
		if !ok && isDigit {
			break
		}
		if isDigit && (ss == "-" || ss == "+") {
			break
		}
		if ok && !isDigit {
			isDigit = true
		}
		if ok && ss == "-" {
			isNegative = true
			continue
		}
		if ok && ss == "+" {
			isNegative = false
			continue
		}
		b.WriteString(ss)
	}
	ss := b.String()
	if isNegative && ss != "" {
		ss = "-" + b.String()
	}
	if ss == "" {
		ss = "0"
	}
	r, err := strconv.Atoi(ss)
	if err != nil {
		if strings.HasPrefix(ss, "-") {
			return minx
		} else {
			return maxx
		}
	}
	if r < minx {
		return minx
	}
	if r > maxx {
		return maxx
	}
	return r
}
