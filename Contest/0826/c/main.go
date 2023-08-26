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

// 深さ優先探索
/*
graph: このmapは、各町（town）から到達可能な他の町への道（road）を格納します。キーは町（整数）で、値はTown構造体のスライスです。
current: 現在訪れている町。
distance: 現在の町に到達するためにかかった距離。
maxDistance: 最大距離。この変数はコード内で明示的には表示されていませんが、グローバル変数または関数の外部で定義されている可能性が高いです。
visited: 各町が訪れられたかどうかを格納するmapまたは配列。このコードでは表示されていませんが、おそらくグローバル変数か関数の外部で定義されているでしょう。
*/
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
	// 再帰が戻ってきたら、現在の町（current）の訪問済みマークを解除します（これにより、他のパスでこの町を再訪することが可能になります）。
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
