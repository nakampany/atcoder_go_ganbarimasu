package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var s string
	fmt.Scan(&s)
	x := make(map[string]int)
	for i := 1; i <= n; i++ {
		var value int
		fmt.Scan(&value)
		x[strconv.Itoa(i)] = value
	}

	fmt.Println(x)

	// 巡回シフト 各要素を1つ右に動かし、最後の要素を最初に移動する
	// 1 2 3 1 2 2 1 2 apzbqrcs
	// 1回目
	// map[1:1 4:1 7:1]
	// cpzaqrbs
	// 2回目
	// map[2:2 5:2 6:2 8:2]
	// cszapqbr
	// 3回目
	// map[3:3]
	// cszapqbr
}
