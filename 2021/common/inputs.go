package common

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadTestInput(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func NumsFromString(input string) []int {
	var out []int
	for _, l := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(l)
		out = append(out, n)
	}
	return out
}
