package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}

	d := make(map[int][]int)
	for i, a := range A {
		d[a] = append(d[a], i)
	}

	ans := make([]int, N)
	s := 0
	var keys []int
	for k := range d {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, v := range keys {
		for _, i := range d[v] {
			ans[i] = s
		}
		s += v * len(d[v])
	}

	for _, v := range ans {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
