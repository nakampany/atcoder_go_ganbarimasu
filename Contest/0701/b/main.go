package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	price := make([]int, m+1)
	for i := 0; i <= m; i++ {
		fmt.Scan(&price[i])
	}

	value := 0

	for i := 0; i < n; i++ {
		found := false
		for j := 0; j < m; j++ {
			if a[i] == b[j] {
				value += price[j+1]
				found = true
				break
			}
		}
		if !found {
			value += price[0]
		}
	}

	fmt.Println(value)
}
