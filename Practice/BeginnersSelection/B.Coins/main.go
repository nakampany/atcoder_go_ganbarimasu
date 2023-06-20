package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var x int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	ans := 0

	x = x - 500*a - 100*b - 50*c

	if x == 0 {
		ans++
	} else {
		for i := 0; i <= a; i++ {
			for j := 0; j <= b; j++ {
				for k := 0; k <= c; k++ {
					if 500*i+100*j+50*k == x {
						ans++
					}
				}
			}
		}
	}

	fmt.Println(ans)
}
