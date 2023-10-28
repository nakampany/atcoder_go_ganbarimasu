package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for {
		hundreds := N / 100
		tens := (N / 10) % 10
		ones := N % 10

		if hundreds*tens == ones {
			fmt.Println(N)
			break
		}

		N++
	}
}
