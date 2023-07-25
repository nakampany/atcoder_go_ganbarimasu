package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	ab, a, b, ba := 0, 0, 0, 0
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		sc.Scan()
		s := sc.Text()
		l := len(s)
		if s[l-1] == 'A' {
			a++
		}
		if s[0] == 'B' {
			b++
		}
		if s[0] == 'B' && s[l-1] == 'A' {
			ba++
		}
		ab += strings.Count(s, "AB")
	}
	if a == ba && b == ba {
		fmt.Println(ab + max(0, ba-1))
	} else {
		fmt.Println(ab + min(a, b))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
