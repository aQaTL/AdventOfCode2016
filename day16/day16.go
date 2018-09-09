package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	input := "10011111011011001"
	length := 272

	wg := sync.WaitGroup{}
	calc := func(input string, length int) {
		fmt.Println(Part1(input, length))
		wg.Done()
	}

	wg.Add(2)
	go calc(input, length)
	go calc(input, 35651584)
	wg.Wait()
}

func Part1(input string, length int) string {
	a := make([]bool, len(input), length*2)
	for i := range input {
		if input[i] == '1' {
			a[i] = true
		} else {
			a[i] = false
		}
	}

	for len(a) < length {
		a = append(a, false)
		for i := len(a) - 2; i >= 0; i-- {
			a = append(a, !a[i])
		}
	}

	checksum := make([]bool, length)
	copy(checksum, a)

	for len(checksum)%2 == 0 {
		newChecksum := make([]bool, 0, len(checksum)/2)
		for i := 0; i < len(checksum)-1; i += 2 {
			newChecksum = append(newChecksum, checksum[i] == checksum[i+1])
		}
		checksum = newChecksum
	}

	builder := strings.Builder{}
	builder.Grow(len(checksum))
	for _, b := range checksum {
		if b {
			builder.WriteByte('1')
		} else {
			builder.WriteByte('0')
		}
	}

	return builder.String()
}
