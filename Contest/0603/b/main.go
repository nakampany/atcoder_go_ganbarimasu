package main

import (
	"fmt"
	"math"
)

const NUMBER = 1000

func main() {
	var n int
	fmt.Scan(&n)

	if n < NUMBER {
		fmt.Println(n)
	} else if NUMBER <= n && n < NUMBER*10 {
		n = int(math.Floor(float64(n)/10)) * 10
		fmt.Println(n)
	} else if NUMBER*10 <= n && n < NUMBER*100 {
		n = int(math.Floor(float64(n)/100)) * 100
		fmt.Println(n)
	} else if NUMBER*100 <= n && n < NUMBER*1000 {
		n = int(math.Floor(float64(n)/1000)) * 1000
		fmt.Println(n)
	} else if NUMBER*1000 <= n && n < NUMBER*10000 {
		n = int(math.Floor(float64(n)/10000)) * 10000
		fmt.Println(n)
	} else if NUMBER*10000 <= n && n < NUMBER*100000 {
		n = int(math.Floor(float64(n)/100000)) * 100000
		fmt.Println(n)
	} else if NUMBER*100000 <= n && n < NUMBER*1000000 {
		n = int(math.Floor(float64(n)/1000000)) * 1000000
		fmt.Println(n)
	}
}
