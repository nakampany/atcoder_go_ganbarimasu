package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	card := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&card[i])
	}

	fmt.Println(card)

	sort.Slice(card, func(i, j int) bool { return card[i] > card[j] })

	fmt.Println(card)

	var alice, bob int

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println("alice")
			fmt.Println(card[i])
			alice += card[i]
		} else {
			fmt.Println("bob")
			fmt.Println(card[i])
			bob += card[i]
		}
	}

	fmt.Println(alice - bob)
}
