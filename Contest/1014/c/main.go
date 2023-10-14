package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	var x string
	fmt.Scan(&n, &x)
	A := make([]string, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		scanner.Scan()
		A[i] = scanner.Text()
	}

	ans := make([]int, 0)

	for i, v := range A {
		// Xは、T と等しい。
		if v == x {
			ans = append(ans, i)
			continue
		}
		// Xは、Tのいずれか1つの位置（先頭と末尾も含む）に英小文字を 1つ挿入して得られる文字列である。
		if len(v) == len(x)-1 {
			for j := 0; j < len(v)+1; j++ {
				if v[:j] == x[:j] && v[j:] == x[j+1:] {
					ans = append(ans, i)
					continue
				}
			}
		}
		// Xは、T からある 1 文字を削除して得られる文字列である。
		if len(v) == len(x)+1 {
			for j := 0; j < len(v); j++ {
				if v[:j] == x[:j] && v[j+1:] == x[j:] {
					ans = append(ans, i)
					continue
				}
			}
		}
		// Xは、T のある 1 文字を別の英小文字に変更して得られる文字列である。
		if len(v) == len(x) {
			for j := 0; j < len(v); j++ {
				if v[j] != x[j] && v[:j]+v[j+1:] == x[:j]+x[j+1:] {
					ans = append(ans, i)
				}
			}
		}
	}

	ans = removeDuplicate(ans)
	sort.Ints(ans)

	fmt.Println(len(ans))
	for _, v := range ans {
		fmt.Printf("%d ", v+1)
	}
	fmt.Println()
}

func removeDuplicate(args []int) []int {
	results := make([]int, 0)
	encountered := map[int]bool{}
	for _, arg := range args {
		if !encountered[arg] {
			encountered[arg] = true
			results = append(results, arg)
		}
	}
	return results
}
