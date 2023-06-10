package main

import (
	"fmt"
)

func main() {
	var p, q byte
	fmt.Scan(&p, &q)

	p -= 'A'
	q -= 'A'

	if p < q {
		p, q = q, p
	}

	a := map[byte]int{'A': 0, 'B': 3, 'C': 4, 'D': 8, 'E': 9, 'F': 14, 'G': 23}

	fmt.Println(a[q] - a[p])

}
