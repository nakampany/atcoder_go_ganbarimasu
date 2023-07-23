package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	// 周期性を考える
	// 1 2 3 1
	// Kが偶数の時、A-B
	// Kが奇数の時、B-A

	// K=1の時、(5,4,3) anser: 5-4=1
	// K=2の時、(7,8,9) anser: 7-8=-1
	// K=3の時、(17,16,15) anser: 17-16=1
	// K=4の時、(31,32,33) anser: 31-32=-1

	if k%2 == 0 {
		fmt.Println(a - b)
	} else {
		fmt.Println(b - a)
	}
}
