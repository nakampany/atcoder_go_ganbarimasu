package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	n := 3
	points := []Point{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		nums := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		points = append(points, Point{x, y})
	}

	fmt.Println(points)
}
