package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	type kv struct {
		Key   int
		Value float64
	}
	var ans []kv

	for i := 0; i < n; i++ {
		ans = append(ans, kv{i, float64(a[i][0]) / float64(a[i][0]+a[i][1])})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Value > ans[j].Value
	})

	var output []int
	for _, v := range ans {
		output = append(output, v.Key+1)
	}
	for _, num := range output {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
