package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scan(&n)

	if len(n) < 2 {
		fmt.Println("Yes")
		return
	}

	a := true
	for i := 0; i < len(n)-1; i++ {
		if n[i] <= n[i+1] {
			a = false
			break
		}
	}

	if a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
