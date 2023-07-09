package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		var s string
		fmt.Scan(&s)
		for j := 0; j < n; j++ {
			a[i][j], _ = strconv.Atoi(string(s[j]))
		}
	}

	// 一時的な配列を作成し、その中にシフトした値を格納します。
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 上辺
			if i == 0 {
				if j == n-1 {
					b[i][j] = a[n-1][0]
				} else {
					b[i][j] = a[i][j+1]
				}
				// 右辺
			} else if j == n-1 {
				if i == n-1 {
					b[i][j] = a[i][0]
				} else {
					b[i][j] = a[i+1][j]
				}
				// 下辺
			} else if i == n-1 {
				if j == 0 {
					b[i][j] = a[i][0]
				} else {
					b[i][j] = a[i][j-1]
				}
				// 左辺
			} else if j == 0 {
				if i == 0 {
					b[i][j] = a[i][n-1]
				} else {
					b[i][j] = a[i-1][j]
				}
			} else {
				// 内部の値は変更しない
				b[i][j] = a[i][j]
			}
		}
	}

	for _, row := range b {
		var str string
		for _, num := range row {
			str += strconv.Itoa(num)
		}
		fmt.Println(str)
	}
}
