package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	edges := make([][]int, n+1)
	for i := 0; i < m; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	// マップのキーは整数（頂点のID）, 値は整数のスライス（その頂点に隣接している他の頂点のIDのリスト）
	colors := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sc.Scan()
		colors[i], _ = strconv.Atoi(sc.Text())
	}
	// fmt.Println(edges)
	// 1: 2, 2: 1 3, 3: 2
	// [[] [2] [1 3] [2]]

	queries := make([][]int, q)
	for i := 0; i < q; i++ {
		query := make([]int, 2)
		sc.Scan()
		query[0], _ = strconv.Atoi(sc.Text())
		sc.Scan()
		query[1], _ = strconv.Atoi(sc.Text())
		if query[0] == 2 {
			sc.Scan()
			newColor, _ := strconv.Atoi(sc.Text())
			query = append(query, newColor)
		}
		queries[i] = query
	}

	fmt.Println(queries)
	// [[1 2] [2 1 20] [1 1]]

	for _, query := range queries {
		node := query[1]
		nodeColor := colors[node]
		fmt.Println(nodeColor)
		if query[0] == 1 {
			// nodeに隣接するすべての頂点の色をnodeColorにする
			for _, u := range edges[node] {
				colors[u] = nodeColor
			}
		} else if query[0] == 2 {
			// nodeの色をnewColorにする
			newColor := query[2]
			colors[node] = newColor
		}
	}
	//	node := query[1]という行で、処理対象のノード（頂点）の番号が取得されます。
	// nodeColor := colors[node]という行で、該当ノードの現在の色が取得されます。
	// 	クエリの種類が1の場合（if query[0] == 1）：
	// colors[u] = nodeColorという行で、各隣接ノード（u）の色を、対象のノードの色（nodeColor）に変更します。

	// クエリの種類が2の場合（else if query[0] == 2）：
	// newColor := query[2]で、新たな色を取得します。
	// そして、colors[node] = newColorで、対象のノードの色を新たな色に変更します。
}
