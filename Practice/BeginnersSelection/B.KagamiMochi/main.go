package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	card := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&card[i])
	}

}
