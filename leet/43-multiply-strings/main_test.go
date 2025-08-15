package main_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		num1   string
		num2   string
		expect string
	}{
		/*
		*/
			{num1: "2", num2: "3", expect: "6"},
			{num1: "123", num2: "456", expect: "56088"},
		{num1: "9133", num2: "0", expect: "0"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("test"), func(t *testing.T) {
			r := multiply(tt.num1, tt.num2)
			if r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	r := make([]byte, l1 + l2)
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			m := (num1[i] - '0') * (num2[j] - '0')
			r[i+j+1] += m
			if r[i+j+1] >= 10 {
				r[i+j] += r[i+j+1] / 10
				r[i+j+1] %= 10
			}
		}
	}
	if r[0] == 0 {
		r = r[1:]
	}
	for i := range len(r) {
		r[i] += '0'
	}
	return string(r)
}

func multiply2(num1 string, num2 string) string {
	arr := make([]string, 0, len(min(num1, num2)))
	for i := len(num1) - 1; i >= 0; i-- {
		var sb strings.Builder
		var t int
		for range len(num1) - 1 - i {
			sb.WriteString("0")
		}
		for j := len(num2) - 1; j >= 0; j-- {
			ii, _ := strconv.Atoi(string(num1[i]))
			jj, _ := strconv.Atoi(string(num2[j]))
			m := ii * jj
			s := m / 10
			a := m % 10
			at := (a + t) % 10
			t = s + ((a + t) / 10)
			sb.WriteString(strconv.Itoa(at))
		}
		if t > 0 {
			sb.WriteString(strconv.Itoa(t))
		}
		arr = append(arr, reverse(sb.String()))
	}
	return checkZero(sum(arr))
}

func checkZero(s string) string {
	for i := range len(s) {
		if string(s[i]) != "0" {
			return s
		}
	}
	return "0"
}

func reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func sum(ss []string) string {
	r := ss[0]
	for i := 1; i < len(ss); i++ {
		j := len(r) - 1
		k := len(ss[i]) - 1
		var t int
		var sb strings.Builder
		for j >= 0 || k >= 0 {
			jj := 0
			if j >= 0 {
				jj, _ = strconv.Atoi(string(r[j]))
			}
			kk := 0
			if k >= 0 {
				kk, _ = strconv.Atoi(string(ss[i][k]))
			}
			p := jj + kk + t
			t = p / 10
			pt := p % 10
			sb.WriteString(strconv.Itoa(pt))

			j--
			k--
		}
		if t > 0 {
			sb.WriteString(strconv.Itoa(t))
		}
		r = reverse(sb.String())
	}
	return r
}
