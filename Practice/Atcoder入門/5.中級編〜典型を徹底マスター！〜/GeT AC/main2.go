package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	var s string
	fmt.Scan(&s)

	acc := make([]int, n+1)
	for i := 1; i < n; i++ {
		if s[i-1] == 'A' && s[i] == 'C' {
			acc[i] = acc[i-1] + 1
		} else {
			acc[i] = acc[i-1]
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(acc[r-1] - acc[l-1])
	}
}
