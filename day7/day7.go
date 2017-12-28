package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	re := regexp.MustCompile("\\[[^\\[-\\]]*\\]")

	tlsMatches := 0
tlsSearch:
	for _, line := range input {
		hypernetSequences := re.FindAllString(line, -1)
		for _, hypernet := range hypernetSequences {
			for i := 0; i < len(hypernet)-3; i++ {
				if hypernet[i] != hypernet[i+1] &&
					hypernet[i] == hypernet[i+3] &&
					hypernet[i+1] == hypernet[i+2] {
					continue tlsSearch
				}
			}
		}
		supernetSequences := re.Split(line, -1)
		for _, supernet := range supernetSequences {
			for i := 0; i < len(supernet)-3; i++ {
				if supernet[i] != supernet[i+1] &&
					supernet[i] == supernet[i+3] &&
					supernet[i+1] == supernet[i+2] {
					tlsMatches++
					continue tlsSearch
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", tlsMatches)

	sslMatches := 0
sslSearch:
	for _, line := range input {
		hypernetSequences := re.FindAllString(line, -1)
		supernetSequences := re.Split(line, -1)

		abas := make([][3]uint8, 0)
		for _, supernet := range supernetSequences {
			for i := 0; i < len(supernet)-2; i++ {
				if supernet[i] != supernet[i+1] && supernet[i] == supernet[i+2] {
					aba := [3]uint8{}
					for j := 0; j < 3; j++ {
						aba[j] = supernet[i+j]
					}
					abas = append(abas, aba)
				}
			}
		}
		if len(abas) == 0 {
			continue sslSearch
		}

		for _, aba := range abas {
			for _, hypernet := range hypernetSequences {
				for i := 0; i < len(hypernet)-2; i++ {
					if hypernet[i] == aba[1] && hypernet[i+2] == aba[1] && hypernet[i+1] == aba[0] {
						sslMatches++
						continue sslSearch
					}
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", sslMatches)
}
