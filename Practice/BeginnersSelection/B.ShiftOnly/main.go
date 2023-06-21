package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var count int

	for {
		for i := 0; i < n; i++ {
			if a[i]%2 == 1 {
				return
			}
			a[i] /= 2
		}
		count++
	}

}
