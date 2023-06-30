package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings.ReplaceAll(s, "eraser", "")
	s = strings.ReplaceAll(s, "erase", "")
	s = strings.ReplaceAll(s, "dreamer", "")
	s = strings.ReplaceAll(s, "dream", "")

	if s == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
