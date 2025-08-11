package main_test

import (
	"strings"
	"strconv"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct{
		n int
		expect string
	}{
		{n: 4, expect: "1211"},
		{n: 1, expect: "1"},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			r := countAndSay(tt.n)
			if r != tt.expect {
				t.Errorf("expect: %v, actual: %v", tt.expect, r)
			}
		})
	}
}
var re string = ""

func countAndSay(n int) string {
	if n == 1 {
		re = "1"
		return re
	} else {
		countAndSay(n - 1)
		var sb strings.Builder
		cnt := 1
		v := string(re[0])
		for i := 1; i < len(re); i++ {
			if string(re[i]) != v && v != "" {
				sb.WriteString(strconv.Itoa(cnt))
				sb.WriteString(string(v))
				cnt = 1
				v = string(re[i])
			} else {
				cnt += 1
			}
		}
		if cnt > 0 {
			sb.WriteString(strconv.Itoa(cnt))
			sb.WriteString(string(v))
		}
		re = sb.String()
	}
	return re
}
