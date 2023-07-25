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
	var a []int64
	for _, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		a = append(a, int64(num))
	}
	numMinus := 0
	mi := int64(1 << 60)
	sum := int64(0)
	for i := 0; i < n; i++ {
		if a[i] < 0 {
			numMinus++
		}
		chmin(&mi, abs(a[i]))
		sum += abs(a[i])
	}
	if numMinus%2 == 0 {
		fmt.Println(sum)
	} else {
		fmt.Println(sum - mi*2)
	}
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func chmin(a *int64, b int64) {
	if *a > b {
		*a = b
	}
}
