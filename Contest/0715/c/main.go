package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		a[i] = sortString(a[i])
	}

	count := make(map[string]bool)
	for _, i := range a {
		count[i] = true
	}

	fmt.Println(count)

	// total := len(count) * (len(count) - 1) / 2
	// fmt.Println(total)
}

func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
