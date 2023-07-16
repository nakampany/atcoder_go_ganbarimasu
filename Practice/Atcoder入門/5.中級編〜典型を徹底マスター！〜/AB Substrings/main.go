package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	var 10000_y int
	var 5000_y int
	var 1000_y int

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k = n-i-j
			if (10000*i+5000j+1000*k == y) {
				fmt.Println(i, j, k)
				break
			}
		}
	}
	fmt.Println(-1, -1, -1)



}
