package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

type Words struct {
	lower string
	big   string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s) // 改行を削除

	words := []Words{}
	buf := strings.Builder{} // 文字列連結用のBuilder
	big_count := 0

	for _, c := range s {
		buf.WriteRune(c) // 文字列連結
		if unicode.IsUpper(c) {
			big_count++
			if big_count%2 == 0 {
				word := buf.String()
				lower := strings.ToLower(word)
				words = append(words, Words{lower, word})
				buf.Reset() // Builderをリセット
			}
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].lower < words[j].lower
	})

	buf.Reset() // 最終結果用のBuilderをリセット
	for _, p := range words {
		buf.WriteString(p.big)
	}
	fmt.Println(buf.String())
}
