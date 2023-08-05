package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	x := a[0]
	ans := 0

	for i := 1; i < n; i++ {
		if x <= a[i] {
			if ans < (a[i]+1)-x {
				ans = (a[i] + 1) - x
			}
		}
	}

	fmt.Println(ans)
}
