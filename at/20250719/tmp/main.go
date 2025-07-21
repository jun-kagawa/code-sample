package main

import (
	"fmt"
)

func main () {
	var n = 5
	for i := 0; i < (1 << n); i++ {
		fmt.Printf("%b\n", i)
		var sum int
		for j := 0; j < n; j++ {
			if (i >> j) & 1 == 1 {
				sum += (1 << j)
			}
		}
		fmt.Println("sum", sum)
	}
}

