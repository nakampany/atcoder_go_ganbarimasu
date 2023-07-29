package main

import (
	"fmt"
)

func main() {
	var n, m int
	var s []string
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scanf("%s", &tmp)
		s = append(s, tmp)
	}

	for i := 0; i < (n - 8); i++ {
		for j := 0; j < (m - 8); j++ {
			var flag bool = true
			// 左上
			if s[i][j] != '#' || s[i][j+1] != '#' || s[i][j+2] != '#' || s[i][j+3] != '.' {
				flag = false
			}
			if s[i+1][j] != '#' || s[i+1][j+1] != '#' || s[i+1][j+2] != '#' || s[i+1][j+3] != '.' {
				flag = false
			}
			if s[i+2][j] != '#' || s[i+2][j+1] != '#' || s[i+2][j+2] != '#' || s[i+2][j+3] != '.' {
				flag = false
			}

			if s[i+3][j] != '.' || s[i+3][j+1] != '.' || s[i+3][j+2] != '.' || s[i+3][j+3] != '.' {
				flag = false
			}

			// 右下
			if s[i+5][j+5] != '.' || s[i+5][j+6] != '.' || s[i+5][j+7] != '.' || s[i+5][j+8] != '.' {
				flag = false
			}
			if s[i+6][j+5] != '.' || s[i+6][j+6] != '#' || s[i+6][j+7] != '#' || s[i+6][j+8] != '#' {
				flag = false
			}
			if s[i+7][j+5] != '.' || s[i+7][j+6] != '#' || s[i+7][j+7] != '#' || s[i+7][j+8] != '#' {
				flag = false
			}

			if s[i+8][j+5] != '.' || s[i+8][j+6] != '#' || s[i+8][j+7] != '#' || s[i+8][j+8] != '#' {
				flag = false
			}

			if flag {
				fmt.Printf("%d %d\n", i+1, j+1)
			}
		}
	}
}
