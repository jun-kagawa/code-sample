package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(a, n int) int {
	var total int
	for digitNum := 1; ; {
		diff := n / a
		total += (n % a) * digitNum
		if diff == 0 {
			break
		}
		n = diff
		digitNum *= 10
	}
	return total
}

func palindromic(a int) bool {
	s := strconv.Itoa(a)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	var a int
	fmt.Scan(&a)
	var n int
	fmt.Scan(&n)
	ns := strconv.Itoa(n)
	d := len(ns) / 2
	if len(ns)%2 == 1 {
		d += 1
	}
	da, _ := strconv.Atoi(ns[0 : d+1])
	var total int
	for i := 1 * (d - 1); i < da; i++ {
		var sb strings.Builder
		is := strconv.Itoa(i)
		sb.WriteString(is)
		for j := len(is) - 1; j >= 0; j-- {
			sb.WriteByte(is[j])
		}
		num, _ := strconv.Atoi(sb.String())
		fmt.Println("i", i, "num", num)

		if palindromic(convert(a, num)) {
			total += i
		}
	}
	fmt.Println(total)
}

/*
func main() {
	var sb strings.Builder
	var n int
	fmt.Scan(&n)

	var t int
	for i := 0; i < n; i++ {
		var c string
		var l int
		fmt.Scan(&c, &l)

		t += l
		if t > 100 {
			fmt.Println("Too Long")
			return
		}
		for j := 0; j < l; j++ {
			sb.WriteString(c)
		}
	}
	fmt.Println(sb.String())
}
*/

/*
func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)

	var c int
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if x <= l && r <= y {
			c++
		}
	}
	fmt.Println(c)
}
*/
