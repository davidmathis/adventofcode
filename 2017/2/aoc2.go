package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
    "strings"
    "strconv"
)

func main() {
	file, err := os.Open("aoc2.txt")
	if err != nil { log.Fatal(err) }

    total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
            lineStr := scanner.Text()
            biggy, _ := strconv.Atoi(string(lineStr[0]))
            lil, _ := strconv.Atoi(string(lineStr[0]))
            nums := strings.Fields(lineStr)
            for num := range nums {
                    n, _ := strconv.Atoi(nums[num])
                    if n < lil {
                            lil = n
                    }
                    if n > biggy {
                            biggy = n
                    }
            }
            total += biggy - lil
    }
    fmt.Println("Total: ", total)
}
