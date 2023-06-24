package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	found := false

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := a[i]
			l := a[j]
			kaibun := k + l
			if i == j {
				continue
			}

			if isPalindrome(kaibun) {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isPalindrome(kaibun string) bool {
	for i := 0; i < len(kaibun)/2; i++ {
		if kaibun[i] != kaibun[len(kaibun)-1-i] {
			return false
		}
	}
	return true
}
