package main

import (
	"fmt"
)

func main() {

	var s string
	fmt.Scan(&s)

	if s == "ACE" || s == "BDF" || s == "CEG" || s == "DFA" || s == "EGB" || s == "FAC" || s == "GBD" {
		s = "Yes"
	} else {
		s = "No"
	}

	fmt.Println(s)

}
