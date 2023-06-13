package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	b := false
	count := 0

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for b == false {
		c1 := 0
		for i := range a {
			if a[i]%2 == 0 {
				a[1] /= 2
				c1++
			}
		}
		if c1 == n {
			b = true
		} else {
			count++
		}
	}
	fmt.Println(count)
}
