package main

import (
	"fmt"
)

func main() {
	var n int
	var a, b int

	fmt.Scan(&n, &a, &b)

	ans := 0

	for i := 1; i <= n; i++ {
		num := i
		total := 0

		for num > 0 {
			temp := num % 10
			total += temp
			num /= 10
		}

		if a <= total && total <= b {
			ans += i
		}
	}

	fmt.Println(ans)

}
