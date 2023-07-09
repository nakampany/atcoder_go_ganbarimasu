package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	type kv struct {
		Key   int
		Value int
	}
	var a []kv
	for i := 0; i <= n; i++ {
		fmt.Scan(&a[i].Key, &a[i].Value)
	}
	fmt.Println(a)
}
