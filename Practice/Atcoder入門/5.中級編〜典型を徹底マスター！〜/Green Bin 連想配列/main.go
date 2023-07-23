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
	counts := make(map[string]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		a[i] = sortString(a[i])
		counts[a[i]]++
	}

	total := 0
	for _, count := range counts {
		if count > 1 {
			total += count * (count - 1) / 2
		}
	}
	fmt.Println(total)
}

// strings.Split(s, ""): "hello" を ["h", "e", "l", "l", "o"]
// sort.Strings(slice): ["h", "e", "l", "l", "o"] -> ["e", "h", "l", "l", "o"] にソート
// strings.Join(slice, ""): ソートされた文字列のスライスを一つの文字列に結合("ehllo")
func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
