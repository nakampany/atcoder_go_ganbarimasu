package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	person := make(map[string]int)

	for i := 0; i < n; i++ {
		var name string
		var age int
		fmt.Scan(&name, &age)
		person[name] = age
	}

	sortedPersons := make([]string, 0, n)
	for name := range person {
		sortedPersons = append(sortedPersons, name)
	}
	sort.Slice(sortedPersons, func(i, j int) bool {
		return person[sortedPersons[i]] < person[sortedPersons[j]]
	})

	startName := sortedPersons[0]
	var startIndex int
	for i, name := range sortedPersons {
		if name == startName {
			startIndex = i
			break
		}
	}

	outputPersons := make([]string, n)
	for i := 0; i < n; i++ {
		outputPersons[i] = sortedPersons[(startIndex+i)%n]
	}

	///   わからん!!!!!
}
