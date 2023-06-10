package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n%10 >= 5 {
		if n%5 >= 3 {
			// 89
			fmt.Print(n + 10 - n%10)
		} else {
			// 67
			fmt.Print(n + 5 - n%10)
		}
	} else if n%10 == 0 || n%10 == 5 {
		fmt.Print(n)
	} else {
		if n%5 >= 3 {
			// 34
			fmt.Print(n + 5 - n%5)
		} else {
			// 12
			fmt.Print(n - n%5)
		}
	}
}
