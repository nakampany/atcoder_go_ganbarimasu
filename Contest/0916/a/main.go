package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// AのB乗を求める
	ans1 := 1
	for i := 0; i < b; i++ {
		ans1 *= a
	}
	// BのA乗を求める
	ans2 := 1
	for i := 0; i < a; i++ {
		ans2 *= b
	}
	ans := ans1 + ans2

	fmt.Println(ans)
}
