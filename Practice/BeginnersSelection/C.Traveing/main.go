package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	place := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		place[t] = []int{x, y}
	}

	fmt.Println(place)
}
