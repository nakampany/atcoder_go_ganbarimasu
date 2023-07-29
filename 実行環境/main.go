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
	var n, m int
	fmt.Scan(&n, &m)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputSliceA := strings.Fields(scanner.Text())
	listA := make([]int, n)
	for i, v := range inputSliceA {
		num, _ := strconv.Atoi(v)
		listA[i] = num
	}
	sort.Ints(listA)

	scanner.Scan()
	inputSliceB := strings.Fields(scanner.Text())
	listB := make([]int, m)
	for i, v := range inputSliceB {
		num, _ := strconv.Atoi(v)
		listB[i] = num
	}
	sort.Ints(listB)

	j := make([]int, 0)
	mMap := make(map[int]int)

	for _, a := range listA {
		if _, ok := mMap[a]; !ok {
			j = append(j, a)
			mMap[a] = a
		}
	}
	for _, b := range listB {
		if _, ok := mMap[b+1]; !ok {
			j = append(j, b+1)
			mMap[b+1] = b + 1
		}
	}
	sort.Ints(j)

	for _, num := range j {
		if sort.Search(len(listA), func(i int) bool { return listA[i] >= num }) >= (m - sort.Search(len(listB), func(i int) bool { return listB[i] >= num })) {
			fmt.Print(num)
			break
		}
	}
}
