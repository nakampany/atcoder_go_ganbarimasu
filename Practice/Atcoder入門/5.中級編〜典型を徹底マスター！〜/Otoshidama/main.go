package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	var N_1000 int
	var N_5000 int
	var N_10000 int

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			k := n - i - j
			if 10000*i+5000*j+1000*k == y && k >= 0 {
				N_10000 = i
				N_5000 = j
				N_1000 = k
				fmt.Println(N_10000, N_5000, N_1000)
				return
			}
		}
	}

	fmt.Println("-1 -1 -1")
}
