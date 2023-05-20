package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := a / b

	if a%b != 0 {
		ans++
	}

	fmt.Println(ans)
}
