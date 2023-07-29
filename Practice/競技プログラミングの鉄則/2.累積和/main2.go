package main

import (
	"fmt"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W) // グリッドのサイズを読み取る

	// 2次元配列 a を作成し、要素を読み取る
	a := make([][]int, H)
	for i := range a {
		a[i] = make([]int, W)
		for j := range a[i] {
			fmt.Scan(&a[i][j]) // a の要素を読み取る
		}
	}

	// 2次元の累積和を計算する
	s := make([][]int, H+1)
	for i := range s {
		s[i] = make([]int, W+1)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			s[i+1][j+1] = s[i][j+1] + s[i+1][j] - s[i][j] + a[i][j]
		}
	}

	// クエリを処理する
	var Q int
	fmt.Scan(&Q) // クエリの数を読み取る
	for q := 0; q < Q; q++ {
		var x1, x2, y1, y2 int
		fmt.Scan(&x1, &x2, &y1, &y2)                               // クエリの範囲を読み取る
		fmt.Println(s[x2][y2] - s[x1][y2] - s[x2][y1] + s[x1][y1]) // クエリの範囲内の総和を出力する
	}
}
