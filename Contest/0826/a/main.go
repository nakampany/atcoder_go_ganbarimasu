package main

import (
	"fmt"
)

func main() {
	var n, h, x int
	fmt.Scan(&n, &h, &x)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sum := h
	count := 0

	for i := 0; i < n; i++ {
		sum += a[i]
		if sum >= x {
			return
		}
		count++
	}

	fmt.Println(count)

}
