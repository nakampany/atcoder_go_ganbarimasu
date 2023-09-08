package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, d, k int
	var p int64
	fmt.Scan(&n, &d, &p)

	f := make([]int64, n)
	s := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&f[i])
	}

	sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })

	s[0] = f[0]
	for i := 0; i < n-1; i++ {
		s[i+1] = s[i] + f[i+1]
	}

	k = (n + d - 1) / d
	var ans int64 = p * int64(k)

	for i := 0; i < k; i++ {
		ans = min(ans, s[int64(n)-1-int64(i)*int64(d)]+(p*int64(i)))
	}

	fmt.Println(ans)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
