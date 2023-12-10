package main

import (
	"fmt"
)

func main() {
	var n int
	var r, c string
	fmt.Scan(&n)
	fmt.Scan(&r, &c)

	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	rowCounts := map[rune]int{'A': 0, 'B': 0, 'C': 0}
	colCounts := map[rune]int{'A': 0, 'B': 0, 'C': 0}

	for i := 0; i < n; i++ {
		ri := rune(r[i])
		ci := rune(c[i])
		grid[i][int(ri-'A')] = ri
		grid[int(ci-'A')][i] = ci
		rowCounts[ri]++
		colCounts[ci]++
	}
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
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '.' {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
	for i := 0; i < n; i++ {
		fmt.Println(string(grid[i]))
	}
}
