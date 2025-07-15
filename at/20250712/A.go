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
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func genPalindromic(n int) []int {
	r := make([]int, 0)
	ns := strconv.Itoa(n)
	l := len(ns) + 1
	for i := 1; i < l; i++ {
		if i == 1 {
			for j := 1; j < 10; j++ {
				r = append(r, j);
			}
			continue
		}
		isEven := i % 2 == 1
		d := i / 2
		if !isEven {
			d -= 1
		}
		start := 1
		for ii := 0; ii < d; ii++ {
			start *= 10
		}
		for j := start; j < start * 10; j++ {
			var sb strings.Builder
			js := strconv.Itoa(j)
			sb.WriteString(js)
			start := len(js) - 1
			if isEven {
				start -= 1
			}
			for k := start; k >= 0; k-- {
				sb.WriteByte(js[k])
			}
			jn, _ := strconv.Atoi(sb.String())
			if jn <= n {
				r = append(r, jn)
			}
		}
	}
	return r
}

func main() {
	var a int
	fmt.Scan(&a)
	var n int
	fmt.Scan(&n)
	total := 0
	r := genPalindromic(n)
	for _, rr := range r {
		cc := convert(a, rr)
		if palindromic(cc) {
			total += rr
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
