package main

import (
	"fmt"
)

var N int
var cumSum [][]int

func main() {
	fmt.Scan(&N)
	var Q int
	fmt.Scan(&Q)

	cumSum = make([][]int, N+1)
	for i := range cumSum {
		cumSum[i] = make([]int, N+1)
	}

	S := make([]string, N)
	for i := range S {
		fmt.Scan(&S[i])
		for j := 1; j <= N; j++ {
			cumSum[i+1][j] = cumSum[i][j] + cumSum[i+1][j-1] - cumSum[i][j-1]
			if S[i][j-1] == 'B' {
				cumSum[i+1][j]++
			}
		}
	}

	for i := 0; i < Q; i++ {
		var A, B, C, D int
		fmt.Scan(&A, &B, &C, &D)
		ans := getCount(C, D) - getCount(A-1, D) - getCount(C, B-1) + getCount(A-1, B-1)
		fmt.Println(ans)
	}
}

func getCount(H, W int) int {
	if H <= N && W <= N {
		return cumSum[H][W]
	}
	Hq := H / N
	Hr := H % N
	Wq := W / N
	Wr := W % N

	var ret int
	ret += getCount(N, N) * Hq * Wq
	ret += getCount(Hr, N) * Wq
	ret += getCount(N, Wr) * Hq
	ret += getCount(Hr, Wr)
	return ret
}
