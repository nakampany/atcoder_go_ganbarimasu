package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(split[0])

	P := make([]int, N)
	C := make([]int, N)
	F := make([][]string, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		split = strings.Split(scanner.Text(), " ")

		P[i], _ = strconv.Atoi(split[0])
		C[i], _ = strconv.Atoi(split[1])

		F[i] = split[2:]
	}

	ans := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans = ans || (P[i] >= P[j] && isSuperset(F[j], F[i]) && (P[i] > P[j] || len(F[j]) > len(F[i])))
		}
	}
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isSuperset(set1, set2 []string) bool {
	mapSet1 := make(map[string]bool)
	for _, val := range set1 {
		mapSet1[val] = true
	}
	for _, val := range set2 {
		if !mapSet1[val] {
			return false
		}
	}
	return true
}
