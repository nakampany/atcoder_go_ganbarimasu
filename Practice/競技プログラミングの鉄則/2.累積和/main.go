package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
	}
	// 累積和
	// s[i] = a[0] + a[1] + ... + a[i-1]
	// s[0] = 0
	// s[1] = a[0]
	// s[2] = a[0] + a[1]
	s := make([]int, N+1)
	for i := 0; i < N; i++ {
		s[i+1] = s[i] + a[i]
	}

	// 区間 [left, right) の総和
	var left, right int
	fmt.Scan(&left, &right)
	fmt.Println(s[right] - s[left])
}
