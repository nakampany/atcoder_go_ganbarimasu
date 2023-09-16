package main

import (
	"fmt"
	"math"
)

func main() {
	var M int
	fmt.Scan(&M)

	S := make([]string, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&S[i])
	}

	ans := math.MaxInt64

	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			for k := 0; k < M/2; k++ {
				if i != j && i != k && j != k && S[0][i%M] == S[1][j%M] && S[0][i%M] == S[2][k%M] {
					ans = min(ans, max(i, j, k))
				}
			}
		}
	}

	if ans < math.MaxInt64 {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}

func max(i, j, k int) int {
	if i > j {
		if i > k {
			return i
		} else {
			return k
		}
	} else {
		if j > k {
			return j
		} else {
			return k
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
