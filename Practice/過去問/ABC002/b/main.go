package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)

	var result []string
	god_word := []string{"a", "i", "u", "e", "o"}

	for _, char := range a {
		isGodWord := false
		for _, god := range god_word {
			if string(char) == god {
				isGodWord = true
				break
			}
		}
		if !isGodWord {
			result = append(result, string(char))
		}
	}
	fmt.Println(strings.Join(result, ""))
}
