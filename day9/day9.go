package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)

	decompressed := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			group := ""
			i++
			for input[i] != ')' {
				group += string(input[i])
				i++
			}
			length, repeat := 0, 0
			if _, err := fmt.Sscanf(group, "%dx%d", &length, &repeat); err != nil {
				panic(fmt.Errorf("%v while scannning %s", err, group))
			}

			i += length
			decompressed += length * repeat
		} else {
			decompressed++
		}
	}
	fmt.Println("Part 1:", decompressed)

	for i := len(input)-1; i >= 0; i-- {

	}
}
