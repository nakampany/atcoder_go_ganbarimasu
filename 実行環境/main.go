package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	var s string
	fmt.Scan(&s)

	// 累積和
	acc := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		if s[i] == 'A' && s[i+1] == 'C' {
			acc[i+2] = acc[i+1] + 1
		} else {
			acc[i+2] = acc[i+1]
		}
	}

	// クエリ
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		fmt.Println(acc[r] - acc[l-1])
	}
}
