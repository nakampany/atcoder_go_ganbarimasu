package main

import (
	"fmt"
)

func main() {
	a := make([]int, 8)
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

}
