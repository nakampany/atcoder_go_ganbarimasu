package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if (a + 1) == b {
		if a == 3 || a == 6 || a == 9 {
			fmt.Println("No")
			return
		}
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
