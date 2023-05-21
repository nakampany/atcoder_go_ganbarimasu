package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	t := 0
	a := 0

	for i := 0; i < n; i++ {
		if s[i] == 'T' {
			t++
		} else {
			a++
		}
	}

	if t > a {
		fmt.Println("T")
	} else if a > t {
		fmt.Println("A")
	} else if string(s[len(s)-1]) == "A" {
		fmt.Println("T")
	} else {
		fmt.Println("A")
	}

}
