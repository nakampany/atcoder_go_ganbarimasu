package main

import (
	"fmt"
)

func main() {
	var M int

	fmt.Scan(&M)

	reels := make([]string, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&reels[i])
	}

	fmt.Println(findMinimumTime(M, reels))
}

func findMinimumTime(M int, reels []string) int {
	maxTime := 2 * M * M * M
	for targetTime := 0; targetTime <= maxTime; targetTime++ {
		targetDigit := reels[0][targetTime%M]
		allMatch := true
		for _, reel := range reels {
			if reel[targetTime%M] != targetDigit {
				allMatch = false
				break
			}
		}
		if allMatch {
			return targetTime
		}
	}
	return -1
}
