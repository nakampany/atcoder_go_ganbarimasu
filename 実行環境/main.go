package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	type kv struct {
		Key   int
		Value uint64
	}
	a := make([]kv, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i].Key, &a[i].Value)
	}

	sort.SliceStable(a, func(i, j int) bool {
		return a[i].Key < a[j].Key
	})

	var value uint64
	for _, v := range a {
		value += v.Value
	}
	for i := 0; i < n; i++ {
		if value <= uint64(k) {
			fmt.Println(i + 1)
			break
		}
		value -= a[i].Value
	}
}
