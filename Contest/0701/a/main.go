package main

import (
	"fmt"
)

func main() {
	a := make([]int, 8)
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	n := len(a) - 1
	for i := 0; i < n; i++ {
		if a[i] > a[i+1] {
			fmt.Println("NO")
			return
		}
		if a[i]%25 != 0 {
			fmt.Println("NO")
			return
		}
		if a[i] < 100 || a[i] > 675 {
			fmt.Println("NO")
			return
		}
	}
	if a[n]%25 != 0 || a[n] < 100 || a[n] > 675 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
