package main

import (
	"fmt"
)

func main() {
	var cells [9]int
	for i := 0; i < 9; i++ {
		fmt.Scan(&cells[i])
	}

	row := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Horizontal rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Vertical rows
		{0, 4, 8}, {2, 4, 6}, // Diagonals
	}

	var order [9]int
	for i := range order {
		order[i] = i
	}

	notDisappoint, all := 0, 0

	permutations(order[:], 0, len(order), func(arr []int) {
		all++
		disappointFlag := false
		for _, r := range row {
			a, b, c := r[0], r[1], r[2]

			if cells[a] == cells[b] && arr[a] < arr[c] && arr[b] < arr[c] {
				disappointFlag = true
			} else if cells[a] == cells[c] && arr[a] < arr[b] && arr[c] < arr[b] {
				disappointFlag = true
			} else if cells[b] == cells[c] && arr[b] < arr[a] && arr[c] < arr[a] {
				disappointFlag = true
			}
		}
		if !disappointFlag {
			notDisappoint++
		}
	})

	fmt.Printf("%.10f\n", float64(notDisappoint)/float64(all))
}

func permutations(arr []int, l int, r int, callback func([]int)) {
	if l == r {
		callback(arr)
		return
	}

	for i := l; i < r; i++ {
		arr[l], arr[i] = arr[i], arr[l]
		permutations(arr, l+1, r, callback)
		arr[l], arr[i] = arr[i], arr[l]
	}
}
