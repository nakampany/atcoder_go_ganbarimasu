package main

import "fmt"

func main() {
	var N, S, M, L int
	fmt.Scan(&N, &S, &M, &L)

	ans := 10 ^ 9
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			for k := 0; k < 100; k++ {
				if i+j+k == N {
					if i*S+j*M+k*L < ans {
						ans = i*S + j*M + k*L
					}
				}
			}
		}
	}
	fmt.Println(ans)

}
