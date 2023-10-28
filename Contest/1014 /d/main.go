package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	type Interval struct {
		Start, End int64
	}
	intervals := make([]Interval, n)
	for i := 0; i < n; i++ {
		var t, d int64
		fmt.Scan(&t, &d)
		intervals[i] = Interval{t, t + d}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	fmt.Println(intervals)

	count := 0
	var lastPrintedTime int64 = 0
	for _, interval := range intervals {
		if interval.Start < lastPrintedTime {
			continue
		}
		count++
		lastPrintedTime = interval.End
	}

	fmt.Println(count)
}
