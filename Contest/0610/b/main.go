package main

import (
	"fmt"
)

func main() {
	var p, q string
	fmt.Scan(&p, &q)

	if p < q {
		p, q = q, p
	}

	d := make(map[string]int)
	d["A"] = 0
	d["B"] = 3
	d["C"] = 4
	d["D"] = 8
	d["E"] = 9
	d["F"] = 14
	d["G"] = 23

	D := d[p] - d[q]

	fmt.Println(D)
}
