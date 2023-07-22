package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	var s string
	fmt.Scan(&s)

	for i := 0; i < q; i++ {
		var l1, l2 int
		fmt.Scan(&l1, &l2)

		word := s[l1-1 : l2]
		count := strings.Count(word, "AC")

		fmt.Println(count)
	}

}
