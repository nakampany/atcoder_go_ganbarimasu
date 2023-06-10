package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	fmt.Println(math.Floor((n+2)/5) * 5)
}
