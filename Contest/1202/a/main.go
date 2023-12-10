package main

import (
	"fmt"
)

func main() {
	var M, D int
	fmt.Scan(&M, &D)
	var y, m, d int
	fmt.Scan(&y, &m, &d)

	if d == D && m == M {
		fmt.Println(y+1, 1, 1)
	} else if d == D {
		fmt.Println(y, m+1, 1)
	} else {
		fmt.Println(y, m, d+1)
	}
}
