package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	arr := make([][]int, H)
	for i := range arr {
		arr[i] = make([]int, W)
		for j := range arr[i] {
			fmt.Scan(&arr[i][j])
		}
	}

	words := []string{"s", "n", "u", "k", "e"}

	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = words[arr[i][j]]
		}
		fmt.Println()
	}

}
