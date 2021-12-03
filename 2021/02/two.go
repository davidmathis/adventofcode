package main

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/davidmathis/adventofcode/2021/common"
)

func calculate(aim bool) int {
	input := c.ReadTestInput("02/input.txt")
	x, f, d := 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		l := strings.Split(line, " ")
		if len(l) < 2 {
			continue
		}
		v, _ := strconv.Atoi(l[1])
		switch l[0] {
		case "forward":
			f += v
			if aim {
				x += d * v
			}
		case "down":
			d += v
		case "up":
			d -= v
		default:
			continue
		}
	}
	if aim {
		return f * x
	}
	return f * d
}

func partOne() int {
	return calculate(false)
}

func partTwo() int {
	return calculate(true)
}

func main() {
	fmt.Println("PartOne:", partOne())
	fmt.Println("PartTwo:", partTwo())
}
