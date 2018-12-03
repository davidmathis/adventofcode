package main

import (
	"bufio"
	"fmt"
	"os"
)

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

// CloseMatch true if strings are almost identical
func CloseMatch(a, b string) bool {
	var diffs int
	for i, letter := range a {
		if letter != rune(b[i]) {
			diffs++
		}
		if diffs > 1 {
			return false
		}
	}
	return true
}

// MakeMatch will update char in strings to make them identical
func MakeMatch(a, b string) string {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[:i] + a[i+1:]
		}
	}
	return ""
}

func main() {
	puzzle, err := readInput("2.txt")
	if err != nil {
		panic(err)
	}

	var lastline string

	for i, line := range puzzle {
		if i >= len(puzzle) {
			continue
		}

		otherlines := puzzle[i+1:]

		for _, other := range otherlines {
			if lastline == "" {
				lastline = other
				continue
			}

			// if near match, replace mismatch and return
			if CloseMatch(line, other) {
				fmt.Println(MakeMatch(line, other))
				break
			}

			lastline = other
		}
	}
}
