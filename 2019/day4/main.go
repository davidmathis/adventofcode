package main

import (
	"fmt"
)

func main() {
	c, start, end := 0, 134564, 585159

	for i := start; i <= end; i++ {
		if isValidPassword(i) {
			c++
		}
	}
	fmt.Println(c)
}

func isValidPassword(i int) bool {
	lastDigit := i % 10
	hasMatch := false
	matches := make(map[int]int)
	for i > 0 {
		// Sorted digits check
		if lastDigit < i/10%10 {
			return false
		}
		// Contiguous digits check
		if lastDigit == i/10%10 {
			hasMatch = true
		}
		// Match counter
		if hasMatch {
			matches[lastDigit]++
		}
		i /= 10
		lastDigit = i % 10
	}
	for _, v := range matches {
		if v == 2 {
			return true
		}
	}
	return false
}
