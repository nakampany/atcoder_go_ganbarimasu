package main

import (
	"fmt"
)

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)
	result := m
	count := 0
	for n >= result {
		result += p
		count++
	}
	fmt.Println(count)
}
