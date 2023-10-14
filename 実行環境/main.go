package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var x string
	fmt.Scan(&n, &x)

	validSet := make(map[string]bool)

	validSet[x] = true

	for i := 0; i <= len(x); i++ {
		for c := 'a'; c <= 'z'; c++ {
			validSet[x[:i]+string(c)+x[i:]] = true
		}
	}

	for i := 0; i < len(x); i++ {
		validSet[x[:i]+x[i+1:]] = true
	}

	for i := 0; i < len(x); i++ {
		for c := 'a'; c <= 'z'; c++ {
			if x[i] != byte(c) {
				validSet[x[:i]+string(c)+x[i+1:]] = true
			}
		}
	}

	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if validSet[s] {
			ans = append(ans, i)
		}
	}
	sort.Ints(ans)

	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Printf("%d ", v+1)
	}
	fmt.Println()
}
