package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	d := y - x

	if d >= -3 && d <= 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
