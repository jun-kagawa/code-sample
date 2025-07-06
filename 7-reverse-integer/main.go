package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	if r := reverse(123); r != 321 {
		panic(r)
	}
	if r := reverse(-123); r != -321 {
		panic(r)
	}
	if r := reverse(120); r != 21 {
		panic(r)
	}
	if r := reverse(1534236469); r != 0 {
		panic(r)
	}
}

var max = math.Pow(2, 31) - 1
var min = math.Pow(-2, 31)

func reverse(x int) int {
	var end int
	var b strings.Builder
	if x < 0 {
		end = 1
		b.WriteString("-")
	}
	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= end; i-- {
		b.WriteByte(s[i])
	}
	v, err := strconv.Atoi(b.String())
	if err != nil {
		return 0
	}
	if v > int(max) || v < int(min) {
		return 0
	}
	return v
}
