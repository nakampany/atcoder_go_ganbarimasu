package main

import (
	"fmt"
)

func main() {
	var n int
	var r, c string
	fmt.Scan(&n)
	fmt.Scan(&r, &c)

	// n x n サイズのグリッドを作成し、全てのセルを'.'で初期化
	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// 各行と各列における'A'、'B'、'C'のカウントを記録するマップを初期化
	rowCounts := map[rune]int{'A': 0, 'B': 0, 'C': 0}
	colCounts := map[rune]int{'A': 0, 'B': 0, 'C': 0}

	// 入力された r, c に基づいて初期の配置を行う
	for i := 0; i < n; i++ {
		ri := rune(r[i])
		ci := rune(c[i])
		grid[i][int(ri-'A')] = ri
		grid[int(ci-'A')][i] = ci
		rowCounts[ri]++
		colCounts[ci]++
	}

	// グリッド内の残りのセルを埋める
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '.' {
				for _, char := range "ABC" {
					if rowCounts[char] < n && colCounts[char] < n && grid[i][int(char-'A')] == '.' && grid[int(char-'A')][j] == '.' {
						grid[i][j] = char
						rowCounts[char]++
						colCounts[char]++
						break
					}
				}
			}
		}
	}

	// グリッドが完全に埋まっているか確認
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '.' {
				fmt.Println("No")
				return
			}
		}
	}

	// グリッドが有効なら"Yes"を出力し、グリッドを表示
	fmt.Println("Yes")
	for i := 0; i < n; i++ {
		fmt.Println(string(grid[i]))
	}
}
