package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	ans := make(map[[2]int]bool)

	for i := 0; i < N; i++ {
		var x1, x2, y1, y2 int
		fmt.Scan(&x1, &x2, &y1, &y2)

		for x := x1; x < x2; x++ {
			for y := y1; y < y2; y++ {
				ans[[2]int{x, y}] = true
			}
		}
	}

	fmt.Println(len(ans))
}
