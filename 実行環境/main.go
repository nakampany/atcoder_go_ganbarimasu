package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		fmt.Scan(&a[i][0])

		var num int
		fmt.Scan(&num)
		for j := 1; j <= num; j++ {
			fmt.Scan(&a[i][j])
		}
		for j := num + 1; j < m; j++ {
			a[i][j] = 0
		}
	}
}
