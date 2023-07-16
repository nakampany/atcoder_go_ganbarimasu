package main

import (
	"fmt"
)

func main() {
	var n, p, q int
	fmt.Scan(&n, &p, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	ans := 1 << 20
	if ans > p {
		ans = p
	}

	for i := 0; i < n; i++ {
		if ans > a[i]+q {
			ans = a[i] + q
		}
	}

	fmt.Println(ans)

}
