package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Pattern is fabric dimension
type Pattern struct {
	x, y, w, h int
}

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

// ParsePattern returns pattern data
func ParsePattern(s string) Pattern {
	var p Pattern
	for i, n := range strings.Split(strings.Replace(s, ":", "", -1), " ") {
		switch i {
		case 2:
			p.x, _ = strconv.Atoi(strings.Split(n, ",")[0])
			p.y, _ = strconv.Atoi(strings.Split(n, ",")[1])
		case 3:
			p.w, _ = strconv.Atoi(strings.Split(n, "x")[0])
			p.h, _ = strconv.Atoi(strings.Split(n, "x")[1])
		}
	}
	return p
}

func main() {
	puzzle, err := readInput("3.bak")
	if err != nil {
		panic(err)
	}

	for _, line := range puzzle {
		p := ParsePattern(line)
		fmt.Printf("%d x %d\n", p.w, p.h)
	}
}
