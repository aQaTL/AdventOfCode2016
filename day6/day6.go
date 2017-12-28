package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	indexes := make([]map[rune]int, len(input[0]))
	for i := range indexes {
		indexes[i] = make(map[rune]int)
	}

	for _, line := range input {
		for i, r := range line {
			indexes[i][r]++
		}
	}

	part1Msg := make([]rune, len(indexes))
	for i, occurrencesMap := range indexes {
		r, count := ' ', 0
		for k, v := range occurrencesMap {
			if v > count {
				r = k
				count = v
			}
		}
		part1Msg[i] = r
	}

	part2Msg := make([]rune, len(indexes))
	for i, occurrencesMap := range indexes {
		r, count := ' ', (1<<31)-1
		for k, v := range occurrencesMap {
			if v < count {
				r = k
				count = v
			}
		}
		part2Msg[i] = r
	}

	fmt.Printf("Part 1: %s\nPart 2: %s\n",
		string(part1Msg),
		string(part2Msg))
}
