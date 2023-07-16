package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.SliceStable(a, func(i, j int) bool { return a[i] < a[j] })

	fmt.Println(a)

}
