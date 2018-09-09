package main

import (
	"fmt"
	"github.com/aqatl/fileutils"
)

func main() {
	input := fileutils.MustLoadLines("input.txt", false)
	discs := make([]Disc, len(input))
	for i := range input {
		d := Disc{}
		fmt.Sscanf(input[i], "Disc #%d has %d positions; at time=%d, it is at position %d.",
			&d.Id, &d.Positions, &d.InitTime, &d.InitPosition)
		discs[i] = d
	}

	fmt.Println(part1(discs))

	//Part 2
	discs = append(discs, Disc{
		discs[len(discs)-1].Id + 1,
		11,
		0,
		0,
	})

	fmt.Println(part1(discs))
}

type Disc struct {
	Id           int
	Positions    int
	InitTime     int
	InitPosition int
}

func part1(discs []Disc) int {
timerLoop:
	for i := 0; ; i++ {
		fallPos := (discs[0].InitPosition + i + 1) % discs[0].Positions
		timer := i + 1
		for j := 1; j < len(discs); j++ {
			timer++
			pos := (discs[j].InitPosition + timer) % discs[j].Positions

			if pos != fallPos {
				continue timerLoop
			}
		}
		return i
	}
}
