package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	a, b := 0, 0
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			if getValueZero(a, b, getInput()) == 19690720 {
				fmt.Println("Found it with", a, b)
				fmt.Println("Answer:", 100*a+b)
				break
			}
			b++
		}
		b = 0
		a++
	}
}

func getInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s := "["
	for scanner.Scan() {
		s += scanner.Text()
	}
	s += "]"
	var intSlice []int
	if err := json.Unmarshal([]byte(s), &intSlice); err != nil {
		log.Fatal("invalid input")
	}
	return intSlice
}

func getValueZero(a, b int, intSlice []int) int {
	intSlice[1] = a
	intSlice[2] = b
	for i, v := range intSlice {
		// Control
		if i%4 == 0 && i+3 <= len(intSlice) {
			keyA := intSlice[i+1]
			keyB := intSlice[i+2]
			val := intSlice[i+3]
			switch v {
			case 1:
				intSlice[val] = intSlice[keyA] + intSlice[keyB]
			case 2:
				intSlice[val] = intSlice[keyA] * intSlice[keyB]
			case 99:
				break
			default:
				log.Fatal("invalid code ", v, " at position ", i)
			}
		}
	}
	return intSlice[0]
}
