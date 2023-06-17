package main

import (
	"fmt"
	"sort"
)

func rearrangeSequence(N int, A []int) []int {
	f := make(map[int][]int)
	for i := 1; i <= N; i++ {
		f[i] = []int{}
	}

	for j, num := range A {
		f[num] = append(f[num], j+1)
	}

	keys := make([]int, 0, N)
	for k := range f {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return f[keys[i]][1] < f[keys[j]][1]
	})

	result := make([]int, 0, N)
	for _, key := range keys {
		result = append(result, key)
	}

	return result
}

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, 3*N)
	for i := 0; i < 3*N; i++ {
		fmt.Scan(&A[i])
	}

	r := rearrangeSequence(N, A)

	for _, num := range r {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
