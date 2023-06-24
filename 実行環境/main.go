package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	fmt.Sscan(sc.Text(), &n)
	sc.Scan()
	s = sc.Text()

	stack := make([]int, 0)
	sBytes := []byte(s)
	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] == '(' {
			stack = append(stack, i)
		} else if sBytes[i] == ')' {
			if len(stack) > 0 {
				start := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				sBytes = append(sBytes[:start], sBytes[i+1:]...)
				i = start - 1
			}
		}
	}

	fmt.Println(string(sBytes))
}
