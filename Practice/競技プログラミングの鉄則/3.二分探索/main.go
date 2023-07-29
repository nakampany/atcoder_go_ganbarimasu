package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var a = []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}

// key の値を持つ要素のインデックスを返す (存在しない場合は -1)
func binarySearch(key int) int {
	left, right := 0, len(a)-1 // 配列 a の左端と右端
	for right >= left {
		mid := left + (right-left)/2 // 区間の真ん中
		if a[mid] == key {
			return mid
		} else if a[mid] > key {
			right = mid - 1
		} else if a[mid] < key {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch(51))  // a[4] = 51 (他に a[3], a[5] も)
	fmt.Println(binarySearch(1))   // a[0] = 1
	fmt.Println(binarySearch(910)) // a[9] = 910

	fmt.Println(binarySearch(52))  // -1
	fmt.Println(binarySearch(0))   // -1
	fmt.Println(binarySearch(911)) // -1

	// ----------------------------------------

	fmt.Println("Start Game!")

	// Aさんの年齢の候補を表す区間を、[left, right) と表す
	// Aさんは、left 歳以上 right 歳未満、right 歳自体は候補に含まれないことに注意
	left, right := 20, 36

	reader := bufio.NewReader(os.Stdin) // 入力を読み取るためのreader

	// Aさんの年齢を1つに絞れないうちは繰り返す
	for right-left > 1 {
		mid := left + (right-left)/2 // 区間の真ん中

		// mid 歳以上かを聞いて、回答を yes/no で受け取る
		fmt.Printf("Is the age same or more than %d ? (yes / no)\n", mid)
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans) // 改行文字を除去

		// 回答の応じて、年齢を絞る
		if ans == "yes" {
			left = mid // mid 歳以上なら、mid 歳以上 right 歳以下になるように
		} else {
			right = mid // mid 歳未満なら、left 歳以上 mid 歳未満になるように
		}
	}

	// ズバリ当てる！
	fmt.Printf("The age is %d!\n", left)
}
