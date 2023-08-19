package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	f := make([]int, N)
	s := make([]int, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		f[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		s[i], _ = strconv.Atoi(scanner.Text())
	}

	max := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			x := cal(f[i], s[i], f[j], s[j])
			if x > max {
				max = x
			}
		}
	}

	fmt.Println(max)
}

func cal(fla1, sat1, fla2, sat2 int) int {
	if sat1 < sat2 {
		sat1, sat2 = sat2, sat1
	}
	if fla1 == fla2 {
		return sat1 + sat2/2
	}
	return sat1 + sat2
}
