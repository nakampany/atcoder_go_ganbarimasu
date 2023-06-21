package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	ans := 0

	for i := range a {
		if a[i] == '1' {
			ans++
		}
	}
	fmt.Println(ans)
}
