package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	scanner := bufio.NewScanner(os.Stdin)

	cal := make([][]bool, n)
	for i := range cal {
		cal[i] = make([]bool, d)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j, char := range strings.Split(line, "") {
			if char == "o" {
				cal[i][j] = true
			}
		}
	}

	var count []int

	nRows := len(cal)
	nCols := len(cal[0])

	for j := 0; j < nCols; j++ {
		allTrue := true
		for i := 0; i < nRows; i++ {
			if !cal[i][j] {
				allTrue = false
				break
			}
		}
		if allTrue {
			count = append(count, j)
		}
	}

	var count1 int
	maxCount := 0

	if len(count) == 0 {
		fmt.Println(0)
		return
	}

	for i := 1; i < len(count); i++ {
		if count[i]-count[i-1] == 1 {
			count1++
			if count1 > maxCount {
				maxCount = count1
			}
		} else {
			count1 = 0
		}
	}

	fmt.Println(maxCount + 1)
}
