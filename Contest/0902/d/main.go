package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	Weight, A, B int
}

func main() {
	var N int
	fmt.Scan(&N)

	var edges []Edge

	index := 1
	for line := 0; line < N-1; line++ {
		for i := 0; i <= line; i++ {
			var weight int
			fmt.Scan(&weight)
			edges = append(edges, Edge{Weight: weight, A: index, B: index + line + 1})
		}
		index++
	}

	// Sort edges in descending order by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight > edges[j].Weight
	})

	chosen := make(map[int]bool)
	totalWeight := 0

	// Select edges from highest weight to lowest
	for _, edge := range edges {
		_, aExists := chosen[edge.A]
		_, bExists := chosen[edge.B]

		if !aExists && !bExists {
			totalWeight += edge.Weight
			chosen[edge.A] = true
			chosen[edge.B] = true
		}
	}

	fmt.Println(totalWeight)
}
