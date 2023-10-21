package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputSlice := strings.Fields(scanner.Text())
	var A []uint
	var total uint
	for _, v := range inputSlice {
		num, _ := strconv.ParseUint(v, 10, 64)
		A = append(A, uint(num))
		total += uint(num)
	}

	avg := total / uint(n)
	remainder := total % uint(n)

	var operations uint
	for i := 0; i < n; i++ {
		if uint(i) < uint((n)-int(remainder)) {
			if A[i] > avg {
				operations += A[i] - avg
			} else if A[i] < avg {
				operations += avg - A[i]
			}
		} else {
			if A[i] > avg+1 {
				operations += A[i] - avg - 1
			} else if A[i] < avg+1 {
				operations += avg + 1 - A[i]
			}
		}
	}

	fmt.Println(operations / 2)
}
