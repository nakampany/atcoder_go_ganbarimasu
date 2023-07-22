package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	fmt.Scan(&s)

	count := 0
	bool1 := false
	bool2 := false
	bool3 := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			count++
			bool1 = true
		}
		if s[i] == 'B' {
			count++
			bool2 = true
		}
		if s[i] == 'C' {
			count++
			bool3 = true
		}
		if bool1 && bool2 && bool3 {
			break
		}
	}
	fmt.Println(count)

}
