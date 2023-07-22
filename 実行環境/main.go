package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if s[j] == '#' {
				a[i][j] = 1
			} else {
				a[i][j] = 0
			}
		}
	}

	for _, row := range a {
		fmt.Println(row)
	}
}
