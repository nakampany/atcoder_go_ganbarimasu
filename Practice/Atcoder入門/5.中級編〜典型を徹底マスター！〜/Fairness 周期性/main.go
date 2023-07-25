package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	for i := 0; i < k; i++ {
		a, b, c = sum_number(a, b, c)
	}

	ans := a - b

	if ans > 1e18 {
		fmt.Println("Unfair")
	} else {
		fmt.Println(ans)
	}
}

func sum_number(a, b, c int) (int, int, int) {
	tmp_a := a
	tmp_b := b
	tmp_c := c
	a = tmp_a + tmp_b
	b = tmp_b + tmp_c
	c = tmp_c + tmp_a

	return a, b, c
}
