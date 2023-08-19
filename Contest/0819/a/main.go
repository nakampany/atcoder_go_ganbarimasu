package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	// [ a, e, i, o, u]
	a := []string{"a", "e", "i", "o", "u"}

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			// remove
			s = s[:i] + s[i+1:]
		}
	}
	fmt.Println(s)

}
