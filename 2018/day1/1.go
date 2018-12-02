package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Freq is map for traking frequency
type Freq map[int]bool

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

func main() {
	puzzle, err := readInput("1.txt")
	if err != nil {
		panic(err)
	}

	o := 0
	t := false
	m := Freq{}
	m[o] = true

	for !t {
		for _, p := range puzzle {
			i, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}

			o += i
			if m[o] {
				fmt.Println("Twice:", o)
				t = true
				break
			}

			m[o] = true
		}
		fmt.Println("Total:", o)
	}
}
