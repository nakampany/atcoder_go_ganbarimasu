package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	s := make([][]byte, h)
	for i := range s {
		fmt.Scan(&s[i])
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				a := [][2]int{
					{i - 1, j},
					{i, j - 1},
					{i, j + 1},
					{i + 1, j},
				}
				cnt := 0

				for _, t := range a {
					k, l := t[0], t[1]
					if 0 <= k && k < h &&
						0 <= l && l < w &&
						s[k][l] == '#' {
						cnt++
					}
				}

				if cnt >= 2 {
					fmt.Println(i+1, j+1)
					return
				}
			}
		}
	}
}
