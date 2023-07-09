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
		Value uint64
	}
	a := make([]kv, n)
	for i := 0; i < n; i++ {
		ans = append(ans, kv{i, math.Floor(float64(a[i][0]) * math.Pow(10, 100) / float64(a[i][0]+a[i][1]))})
	}

	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i].Value > ans[j].Value
	})

	var value uint64
	for _, v := range a {
		value += v.Value
	}
	for i := 0; i < n; i++ {
		if value <= uint64(k) {
			fmt.Println(i + 1)
			break
		}
		value -= a[i].Value
	}
}
