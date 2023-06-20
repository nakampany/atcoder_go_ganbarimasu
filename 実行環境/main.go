package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}

	se := make(map[int]bool)
	for i := 0; i < n; i++ {
		se[t[i]] = true
		fmt.Println(se)
	}
	ans := len(se)
	// 	map[10:true]
	// map[8:true 10:true]
	// map[6:true 8:true 10:true]
	fmt.Println(ans)

}
