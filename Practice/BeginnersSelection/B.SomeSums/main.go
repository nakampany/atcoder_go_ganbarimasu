package main

import "fmt"

func main() {
	num := 1897
	total := 0

	for num > 0 {
		temp := num % 10
		total += temp
		num /= 10
		fmt.Println(temp)
		// fmt.Println(num)
		// fmt.Println(total)
	}

	fmt.Println(total)
}
