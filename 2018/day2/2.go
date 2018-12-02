package main

import (
	"bufio"
	"fmt"
	"os"
)

// Freq is map for tracking frequency
type Freq map[rune]int

// readInput parses puzzle input
func readInput(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// countLetters returns map of letters
func countLetters(s string) Freq {
	m := Freq{}
	for _, letter := range s {
		m[letter]++
	}
	return m
}

// main
func main() {
	puzzle, err := readInput("2.txt")
	if err != nil {
		panic(err)
	}

	var a, b int

	for _, line := range puzzle {
		var double, triple bool
		m := countLetters(line)

		for _, v := range m {
			switch v {
			case 2:
				double = true
			case 3:
				triple = true
			default:
				continue
			}
		}

		if double {
			a++
		}

		if triple {
			b++
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
