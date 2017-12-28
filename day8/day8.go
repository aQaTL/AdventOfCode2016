package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(data)
	}
	input := strings.Split(string(data), "\n")

	const width, height = 50, 6
	screen := [height][width]bool{}

	for _, line := range input {
		cmds := strings.Split(line, " ")
		if cmds[0] == "rect" {
			width, height := 0, 0
			fmt.Sscanf(cmds[1], "%dx%d", &width, &height)
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					screen[i][j] = true
				}
			}
		} else if cmds[1] == "row" {
			idx, shift := 0, 0
			fmt.Sscanf(line, "rotate row y=%d by %d", &idx, &shift)

			newRow := [width]bool{}
			for i := 0; i < width; i++ {
				newRow[(i+shift)%width] = screen[idx][i]
			}
			screen[idx] = newRow
		} else if cmds[1] == "column" {
			idx, shift := 0, 0
			fmt.Sscanf(line, "rotate column x=%d by %d", &idx, &shift)

			newColumn := [height]bool{}
			for i := 0; i < height; i++ {
				newColumn[(i+shift)%height] = screen[i][idx]
			}
			for i, cell := range newColumn {
				screen[i][idx] = cell
			}
		} else {
			panic(fmt.Errorf("non parsable command: %s", line))
		}
	}

	fmt.Println("Part 2:")
	lit := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if screen[i][j] {
				fmt.Print("\u2588")
				lit++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Part 1: %d\n", lit)
}
