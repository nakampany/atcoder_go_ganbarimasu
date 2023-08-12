package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	x := make([]int, n)
	for i := 1; i <= n; i++ {
		x[i] = i
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
	}

	fmt.Println(x)

}
