package main

import "fmt"

func main() {
	a := 45 // 45 = 0b101101
	b := 25 // 25 = 0b011001
	// a|b = 0b111101 = 61
	fmt.Println(a, "OR", b, "=", a|b)
	a = 45 // 45 = 0b101101
	b = 25 // 25 = 0b011001
	// a&b = 0b001001 = 9
	fmt.Println(a, "OR", b, "=", a&b)

	A := 0x2d // 45 = 0b101101
	B := 0x19 // 25 = 0b011001
	fmt.Printf("%d AND %d = %d\n", A, B, A&B)
	fmt.Printf("%08b AND %08b = %08b\n", A, B, A&B)
}
