package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var s string
	fmt.Scan(&s)

	x := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		value, _ := strconv.Atoi(scanner.Text())
		x[strconv.Itoa(i)] = value
	}

	s1 := []rune(s)
	for i := 1; i <= m; i++ {
		var positions []int
		for key, val := range x {
			if val == i {
				position, _ := strconv.Atoi(key)
				positions = append(positions, position-1)
			}
		}
		s1 = rotate(s1, positions)
	}

	fmt.Println(string(s1))
}

func rotate(s []rune, p []int) []rune {
	n := len(p)
	if n == 0 {
		return s
	}

	last := s[p[n-1]]
	for i := n - 1; i > 0; i-- {
		s[p[i]] = s[p[i-1]]
	}
	s[p[0]] = last

	return s
}
