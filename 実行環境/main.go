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

	// マップのキーは整数（頂点のID）, 値は整数のスライス（その頂点に隣接している他の頂点のIDのリスト）
	edges := make([][]int, n+1)

	for i := 0; i < m; i++ {
		sc.Scan()
		u, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}

	fmt.Println(edges)
	// 1: 2, 2: 1 3, 3: 2
	// [[] [2] [1 3] [2]]
	// [[] [] [] [9] [6 17] [] [4 23] [23] [30] [3] [] [13] [] [11] [30] [] [] [4] [24] [] [] [] [] [7 6] [18] [26] [25] [] [] [] [14 8]]

	colors := make([]int, n+1) // colors[0]は使わない
	for i := 1; i <= n; i++ {
		sc.Scan()
		colors[i], _ = strconv.Atoi(sc.Text())
	}

	for i := 0; i < q; i++ {
		sc.Scan()
		query, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		node, _ := strconv.Atoi(sc.Text())
		nodeColor := colors[node]
		fmt.Println(nodeColor) // クエリ1,2ともにここは同じ

		if query == 1 {
			// nodeに隣接するすべての頂点の色をnodeColorにする
			for _, u := range edges[node] {
				colors[u] = nodeColor
			}
		} else if query == 2 {
			// nodeの色をnewColorにする
			sc.Scan()
			newColor, _ := strconv.Atoi(sc.Text())
			colors[node] = newColor
		}
	}
}
