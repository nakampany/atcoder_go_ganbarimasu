package main

import (
	"bufio"
	"fmt"
	"os"
)

func square(a int) int {
	return a * a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	var d int64
	fmt.Scan(&n, &d)
	type Point struct {
		x int
		y int
	}
	points := []Point{}
	for i := 0; i < n; i++ {
		var x, y int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		points = append(points, Point{x, y})
	}
	dists := [][]int64{}
	for i := 0; i < n; i++ {
		dists = append(dists, make([]int64, n))
		for j := 0; j < n; j++ {
			var d int64
			d = int64(square(points[i].x-points[j].x) + square(points[i].y-points[j].y))
			dists[i][j] = d
		}
	}
	q := []int{0}
	infected := make([]bool, n)
	infected[0] = true

	for len(q) >= 1 {
		target := q[0]
		q = append(q[:0], q[1:]...)
		for x, v := range dists[target] {
			if v <= d*d && x != target && !infected[x] {
				q = append(q, x)
				infected[x] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		if infected[i] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
