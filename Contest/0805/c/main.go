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
	var A []int
	for _, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		A = append(A, num)
	}

	sort.Ints(A)

	var operations int
	for A[n-1]-A[0] > 1 {
		A[n-1] -= 1
		A[0] += 1
		sort.Ints(A)
		operations += 1
	}

	fmt.Println(operations)
}
