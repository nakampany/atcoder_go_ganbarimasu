package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	check := false

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			if 10000*i+5000*j+1000*k == y {
				fmt.Println(i, j, k)
				check = true
				return
			}
		}
	}

	if !check {
		fmt.Println(-1, -1, -1)
	}

}
