package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Scan(&m)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&x[i])
	}

	sum := 0
	for i := 0; i < m; i++ {
		sum += x[i]
	}

	result := (sum + 1) / 2

	month := 0
	day := 0
	for m, d := range x {
		day += d
		if day >= result {
			month = m + 1
			break
		}
	}

	day = result - (day - x[month-1])

	fmt.Println(month, day)
}
