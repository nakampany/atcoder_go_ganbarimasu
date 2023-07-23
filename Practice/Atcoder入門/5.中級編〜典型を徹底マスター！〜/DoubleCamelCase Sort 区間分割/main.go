package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)

	type Words struct {
		lower string
		big   string
	}

	words := []Words{}
	buf := ""      // 大文字が出現するまでの文字を一時的に保持するためのバッファ
	big_count := 0 // 大文字が出現した回数を数える変数。

	for _, c := range s {
		buf += string(c)
		if unicode.IsUpper(c) { // 文字が大文字かどうかをチェック
			big_count++ // 大文字の出現回数をカウントアップ
			if big_count%2 == 0 {
				lower := strings.ToLower(buf)
				words = append(words, Words{lower, buf})
				buf = ""
			}
		}
	}
	// [{fish FisH} {dog DoG} {cat CaT} {aa AA} {aaa AaA} {abc AbC}]
	// sortしないといけないから、小文字が必要

	sort.Slice(words, func(i, j int) bool {
		return words[i].lower < words[j].lower
	})

	ans := ""
	for _, p := range words {
		ans += p.big
	}
	fmt.Println(ans)
}
