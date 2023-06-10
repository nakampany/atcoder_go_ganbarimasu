package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	fmt.Print((math.Round(n / 5)) * 5)
}
