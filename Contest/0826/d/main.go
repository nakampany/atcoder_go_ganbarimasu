package main

import (
	"fmt"
)

type Town struct {
	town int
	road int
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	x := make(map[int][]Town)
	for i := 0; i < M; i++ {
		var A, B, C int
		fmt.Scan(&A, &B, &C)
		x[A] = append(x[A], Town{town: B, road: C})
		x[B] = append(x[B], Town{town: A, road: C})
	}

}
