package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	// 二次元配列を用意
	b := make([][]bool, n)
	for i := 0; i < n; i++ {
		// b[i] = [false false ... false]
		b[i] = make([]bool, n)
	}

	// ６✖️６のマス目を想像
	for i := 0; i < m; i++ {
		for j := 0; j < n-1; j++ {
			num1 := a[i][j]
			num2 := a[i][j+1]
			b[num1-1][num2-1] = true
			b[num2-1][num1-1] = true
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if !b[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
