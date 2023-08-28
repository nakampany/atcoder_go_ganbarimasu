package main

import (
	"fmt"
)

type Town struct {
	town int
	road int
}

var maxDistance int
var visited []bool

func dfs(graph map[int][]Town, current int, distance int) {
	if distance > maxDistance {
		maxDistance = distance
	}
	visited[current] = true

	for _, Town := range graph[current] {
		if !visited[Town.town] {
			dfs(graph, Town.town, distance+Town.road)
		}
	}

	visited[current] = false
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

	maxDistance = 0
	visited = make([]bool, N+1)

	for i := 1; i <= N; i++ {
		dfs(x, i, 0)
	}

	fmt.Println(maxDistance)
}
