package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	// Aが巣出て等しいかどうかを確認する
	for i := 0; i < n-1; i++ {
		if A[i] != A[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}
