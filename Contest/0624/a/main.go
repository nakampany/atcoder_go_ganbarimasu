package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 7*n)
	for i := 0; i < 7*n; i++ {
		fmt.Scan(&a[i])
	}

	v := []int{}

	for i := 0; i < n; i++ {
		ans := 0
		for j := 0; j < 7; j++ {
			ans += a[j+7*i]
		}

		v = append(v, ans)
	}

	fmt.Println(strings.Trim(fmt.Sprint(v), "[]"))
}
