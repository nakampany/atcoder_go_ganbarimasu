package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			scanner.Scan()
			val, _ := strconv.Atoi(scanner.Text())
			a[i][j] = val
		}
	}

	type kv struct {
		Key   int
		Value float64
	}
	var ans []kv

	for i := 0; i < n; i++ {
		ans = append(ans, kv{i, math.Floor(float64(a[i][0]) * math.Pow(10, 100) / float64(a[i][0]+a[i][1]))})
	}

	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i].Value > ans[j].Value
	})

	var output []int
	for _, v := range ans {
		output = append(output, v.Key+1)
	}
	for _, num := range output {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
