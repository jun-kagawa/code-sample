package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		var s string
		fmt.Scan(&s)
		dan := "0" + s
		ok := make([]bool, 1<<n)
		ok[0] = true

		for k := 0; k < (1 << n); k++ {
			if !ok[k] {
				continue
			}
			for j := 0; j < n; j++ {
				if (k & (1 << j)) != 0 {
					continue
				}
				next := k | (1 << j)
				if string(dan[next]) == "0" {
					ok[next] = true
				}
			}
		}
		if ok[(1<<n)-1] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func recur(sum int, nums []int, ma map[int]bool, m string) bool {
	fmt.Println("ma", ma)
	if len(nums) == len(ma) {
		return true
	}
	for _, n := range nums {
		if _, ok := ma[n]; ok {
			continue
		}
		s := sum
		s += 1 << (n - 1)
		fmt.Println("s", s)
		if string(m[s]) == "1" {
			continue
		}
		ma[n] = true
		if recur(s, nums, ma, m) {
			return true
		}
	}
	return false
}

func mainB() {
	var s string
	fmt.Scan(&s)

	arr := []int{}
	for i := 0; i < len(s); i++ {
		ss := string(s[i])
		if ss == "#" {
			arr = append(arr, i+1)
		}
	}
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d,%d\n", arr[i], arr[i+1])
	}
}

func mainA() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	var x int
	fmt.Scan(&x)

	for _, a := range arr {
		if a == x {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
