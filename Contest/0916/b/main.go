package main

import (
	"fmt"
)

// 回文かどうか
func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// 最長回文部分文字列を求める
func maxPalindromeSubstring(s string) int {
	n := len(s)
	maxLength := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if isPalindrome(s[i:j]) && j-i > maxLength {
				maxLength = j - i
			}
		}
	}
	return maxLength
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(maxPalindromeSubstring(s))
}
