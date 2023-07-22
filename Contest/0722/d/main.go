package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([][]string, n)
	for i := range a {
		a[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
			s := strings.Split(a[i][j], "")
			if s[0] == "#" {
				a[i][j] = "1"
			} else {
				a[i][j] = "0"
			}
		}
	}

	printMatrix(a)
}

func printMatrix(a [][]string) {
	for _, row := range a {
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}
