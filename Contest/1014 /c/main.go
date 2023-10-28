package main

import (
	"fmt"
)

func main() {

	var h, w int
	fmt.Scan(&h, &w)
	grid := make([][]rune, h)
	visit := make([][]bool, h)

	for i := range grid {
		var s string
		fmt.Scan(&s)
		grid[i] = []rune(s)
		visit[i] = make([]bool, w)
	}

	// 上下左右斜め
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// dfsかな
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= h || y < 0 || y >= w || grid[x][y] == '.' || visit[x][y] {
			return
		}
		visit[x][y] = true
		for i := 0; i < 8; i++ {
			nx, ny := x+dx[i], y+dy[i]
			dfs(nx, ny)
		}
	}

	count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '#' && !visit[i][j] {
				dfs(i, j)
				count++
			}
		}
	}

	fmt.Println(count)
}
