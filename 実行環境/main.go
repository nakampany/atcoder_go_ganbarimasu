package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	ms := make(map[int]int)
	ss := make(map[int]int)

	for i := 0; i < N; i++ {
		scanner.Scan()
		f, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s, _ := strconv.Atoi(scanner.Text())

		if val, exists := ms[f]; exists {
			if s > val {
				ss[f] = val
				ms[f] = s
			} else if s > ss[f] {
				ss[f] = s
			}
		} else {
			ms[f] = s
			ss[f] = 0
		}
	}

	max := 0
	for f1, s1 := range ms {
		for f2, s2 := range ms {
			if f1 != f2 {
				total := s1 + s2
				if total > max {
					max = total
				}
			} else {
				total := s1 + ss[f2]/2
				if total > max {
					max = total
				}
			}
		}
	}

	fmt.Println(max)
}
