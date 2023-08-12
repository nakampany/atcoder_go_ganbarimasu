package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	PI100 := "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"

	if n == 0 {
		fmt.Println("3")
	} else {
		fmt.Println(PI100[:n+2]) // n+2 is because of the digit before the dot and n digits after the dot
	}
}
