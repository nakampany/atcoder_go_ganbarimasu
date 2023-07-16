package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) > 0 {
				start := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				s = s[:start] + s[start+1:i] + s[i+1:]
				i -= 2
			}
		}
	}

	fmt.Println(s)
}
