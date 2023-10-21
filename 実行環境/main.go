package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	x := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		x[i] = make(map[int]bool)
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		x[a][b] = true
		for k, _ := range x[b] {
			x[a][k] = true
		}
		for k, _ := range x {
			if k != a && x[k][a] {
				for l, _ := range x[a] {
					x[k][l] = true
				}
			}
		}
	}

	strongest := -1
	for i := 1; i <= n; i++ {
		if len(x[i]) == n-1 {
			strongest = i
			break
		}
	}

	fmt.Println(strongest)
}
