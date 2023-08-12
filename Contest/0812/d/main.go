package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	var t1, t2 int
	var t3 byte
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t1, &t2, &t3)
	}

}
