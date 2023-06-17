package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	strNumbers := strings.Split(input, " ")

	var numbers []int
	for _, strNum := range strNumbers {
		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}

	var ans float64
	ans = 0

	for i, v := range numbers {
		ans += float64(v) * math.Pow(2, float64(i))
	}

	fmt.Println(int(ans))
}
