package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := 100

	for i := 0; i <= 100; i += 5 {
		if math.Abs(float64(n-ans)) > math.Abs(float64(n-i)) {
			ans = i
		}
	}
	fmt.Println(ans)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	ans := 100

// 	for i := 0; i <= 100; i += 5 {
// 		if abs(n-ans) > abs(n-i) {
// 			ans = i
// 		}
// 	}
// 	fmt.Println(ans)
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
