package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputSlice := strings.Fields(scanner.Text())
	var A []int
	for _, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		A = append(A, num)
	}
	fmt.Println(A)

	var B []int
	for _, v := range inputSlice {
		num, _ := strconv.Atoi(v)
		B = append(B, num)
	}
	fmt.Println(B)

}
