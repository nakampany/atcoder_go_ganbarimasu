package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)

	type Base struct {
		W int
		X int
	}
	bases := make([]Base, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&bases[i].W, &bases[i].X)
	}

	var inter []struct {
		start, end, W int
	}

	for _, base := range bases {
		start := (9 - base.X + 24) % 24
		end := (18 - base.X + 24) % 24
		if start <= end {
			inter = append(inter, struct{ start, end, W int }{start, end, base.W})
		} else {
			inter = append(inter, struct{ start, end, W int }{0, end, base.W})
			inter = append(inter, struct{ start, end, W int }{start, 24, base.W})
		}
	}

	maxCount := 0
	for h := 0; h < 24; h++ {
		count := 0
		for _, interval := range inter {
			if interval.start <= h && h < interval.end {
				count += interval.W
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}

	fmt.Println(maxCount)
}
