package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	inputSlice := strings.Fields(s.Text())
	graph := make([]int, n+1)
	for i, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		graph[i+1] = num
	}

	// fmt.Println(graph) // [0 3 7 4 7 3 3 8 2]

	visited := make([]bool, n+1)
	path := make([]int, 0)

	var cycleStart int
	var dfs func(int)

	// 深さ優先探索
	dfs = func(node int) {
		path = append(path, node)
		if visited[node] {
			for i, v := range path {
				if v == node {
					cycleStart = i
					break
				}
			}
			return
		}
		fmt.Println(visited)
		visited[node] = true
		dfs(graph[node])
	}

	dfs(1)
	count := len(path) - cycleStart - 1
	fmt.Println(count)

	for i, v := range path[cycleStart:] {
		fmt.Print(v, " ")
		if i == count-1 {
			fmt.Println()
			break
		}
	}
}
