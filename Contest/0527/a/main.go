package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b string
	fmt.Scan(&b)
	var c string
	fmt.Scan(&c)

	b_1 := strings.Replace(b, "o", "0", a)
	b_2 := strings.Replace(b_1, "l", "1", a)
	c_1 := strings.Replace(c, "o", "0", a)
	c_2 := strings.Replace(c_1, "l", "1", a)

	if b_2 == c_2 {
		fmt.Println("Yes")

	} else {
		fmt.Println("No")

	}

}
