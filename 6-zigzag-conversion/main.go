package main

import (
	"strings"
)

func main() {
	if r := convert("PAYPALISHIRING", 3); r != "PAHNAPLSIIGYIR" {
		panic(r)
	}
	if r := convert("PAYPALISHIRING", 4); r != "PINALSIGYAHRPI" {
		panic(r)
	}
	if r := convert("A", 1); r != "A" {
		panic(r)
	}
	if r := convert("ABCD", 3); r != "ABDC" {
		panic(r)
	}
}

func convert(s string, numRows int) string {
	var b strings.Builder
	base := numRows*2 - 2
	if base == 0 {
		return s
	}
	for i := range numRows {
		diff := (numRows-i)*2 - 2
		if diff == base {
			diff = 0
		}
		for j := i; j < len(s); j += base {
			b.WriteByte(s[j])
			if diff > 0 && j+diff < len(s) {
				b.WriteByte(s[j+diff])
			}

		}

	}
	return b.String()
}
