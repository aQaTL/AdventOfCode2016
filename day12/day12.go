package main

import "fmt"

func partOne() int {
	a := 1
	b := 1
	c := 0
	d := 26

	for ; d != 0; d-- {
		c = a
		a += b
		b = c
	}

	a += 19 * 14

	return a
}

func partTwo() int {
	a := 1
	b := 1
	c := 0
	d := 26

	d += 7

	for ; d != 0; d-- {
		c = a
		a += b
		b = c
	}

	a += 19 * 14

	return a
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
