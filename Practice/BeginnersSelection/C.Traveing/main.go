package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	place := make(map[int][]int)
	timeKeys := make([]int, n)
	for i := 0; i < n; i++ {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		place[t] = []int{x, y}
		timeKeys[i] = t
	}

	sort.Ints(timeKeys) // Sort the timeKeys

	t0, x0, y0 := 0, 0, 0
	for _, timeKey := range timeKeys {
		v := place[timeKey]

		t := timeKey
		x := v[0]
		y := v[1]

		if (math.Abs(float64(x0-x)) + math.Abs(float64(y0-y))) > float64(t-t0) {
			fmt.Println("No")
			return
		}

		if int(math.Abs(float64(x-x0))+math.Abs(float64(y-y0))-float64(t-t0))%2 != 0 {
			fmt.Println("No")
			return
		}
		t0, x0, y0 = t, x, y
	}

	fmt.Println("Yes")
}
