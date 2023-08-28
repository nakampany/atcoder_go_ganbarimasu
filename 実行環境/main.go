package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	fmt.Scan(&m)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&x[i])
	}

	// ｘを昇順にソート
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	fmt.Println(x)

}
