package main

import (
	"fmt"

	c "github.com/davidmathis/adventofcode/2021/common"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func getNumsFromString() []int {
	input := c.ReadTestInput("01/input.txt")
	return c.NumsFromString(input)
}

func partOne() int {
	nums := getNumsFromString()
	total, last := 0, 0
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] > last {
			total++
		}
		last = nums[i]
	}
	return total
}

func partTwo() int {
	nums := getNumsFromString()
	total, last := 0, 0
	for i := 0; i < len(nums)-2; i++ {
		sum := nums[i] + nums[i+1] + nums[i+2]
		if i != 0 && sum > last {
			total++
		}
		last = sum
	}
	return total
}
