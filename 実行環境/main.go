package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for _, v := range inputSlice {
		num, _ := strconv.ParseUint(v, 10, 32)
		A = append(A, uint(num))
	}

	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })

	var x uint
	for A[n-1]-A[0] > 1 {
		A[n-1] -= 1
		A[0] += 1
		sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
		x += 1
	}

	fmt.Println(x)
}
