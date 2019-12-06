package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuel := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel := (n / 3) - 2
		totalFuel += fuel
		for fuel > 0 {
			fuel = (fuel / 3) - 2
			if fuel > 0 {
				totalFuel += fuel
			}
		}
	}
	fmt.Println(totalFuel)
}
