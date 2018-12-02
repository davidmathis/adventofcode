package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"log"
)

func main() {

	output := 0
	file, err := ioutil.ReadFile("aoc1.txt")
	f := strings.Replace(string(file), "\n", "", -1)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(f); i++ {
		if i == len(f)-1 {
			if f[i] == f[0] {
				num, err := strconv.Atoi(string(f[i]))
				if err != nil {
					log.Fatal(err)
				}
				output += num
			}
		} else {
			if f[i] == f[i+1] {
				num, err := strconv.Atoi(string(f[i]))
				if err != nil {
					log.Fatal(err)
				}
				output += num
			}
		}
	}

	fmt.Println("Answer: ", output)

}
