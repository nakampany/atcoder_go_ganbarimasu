package main

import "fmt"

var n int

func main() {
	fmt.Scan(&n)
	s := make([]string, n)
	a := make([]int, n)

	min := int(2e9)
	minIdx := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &a[i])
		if a[i] < min {
			min = a[i]
			minIdx = i
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(s[(i+minIdx)%n])
	}
}
