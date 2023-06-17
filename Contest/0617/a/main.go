package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var v string
	fmt.Scan(&v)

	ans := make([]string, 100*n)

	for i := 1; i < n+1; i++ {
		ans[2*i-1] = string(v[i-1])
		ans[2*i] = string(v[i-1])
	}

	a := strings.Join(ans[:], "")

	fmt.Println(a)
}
