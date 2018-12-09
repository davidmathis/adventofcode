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

// GenSquare plots the pattern
func GenSquare(x, y, w, h int) map[string]bool {
	coords := make(map[string]bool)
	for row := 1; row <= h; row++ {
		for width := 1; width <= w; width++ {
			pos := strconv.Itoa(width+x) + "," + strconv.Itoa(row+y)
			coords[pos] = true
		}
	}
	return coords
}

func main() {
	puzzle, err := readInput("3.txt")
	if err != nil {
		panic(err)
	}

	m := make(map[string]bool)
	var overlaps int

	for _, line := range puzzle {
		p := ParsePattern(line)
		for k, _ := range GenSquare(p.x, p.y, p.w, p.h) {
			if _, ok := m[k]; ok {
				overlaps++
			} else {
				m[k] = true
			}
		}
	}
	fmt.Println("overlaps:", overlaps)
}
