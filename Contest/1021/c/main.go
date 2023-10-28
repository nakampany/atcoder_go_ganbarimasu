package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)

	maxPresents := 0
	for i := 0; i < n; i++ {
		low, high := i, n
		for low < high {
			mid := (low + high) / 2
			if a[mid] < a[i]+m {
				low = mid + 1
			} else {
				high = mid
			}
		}
		if maxPresents < low-i {
			maxPresents = low - i
		}
	}

	fmt.Println(maxPresents)
}
