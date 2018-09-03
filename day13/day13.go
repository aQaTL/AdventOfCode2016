package main

import (
	"fmt"
	"strconv"
	"github.com/aqatl/fileutils"
)

type Coord struct {
	x, y int
}

func (c Coord) OneD(sideSize int) int {
	return c.y*sideSize + c.x
}

func (c Coord) Around1D(sideSize int) (up, down, right, left int) {
	up = (c.y-1)*sideSize + c.x
	down = (c.y+1)*sideSize + c.x
	left = c.y*sideSize + (c.x - 1)
	right = c.y*sideSize + (c.x + 1)
	return
}

func main() {
	input, _ := strconv.Atoi(fileutils.MustLoadToString("input.txt"))

	srcCoord := Coord{1, 1}
	dstCoord := Coord{31, 39}

	sideSize := 100
	maze := GenMaze(sideSize, input)

	//Part one
	fmt.Println(LeeAlgorithm(maze, sideSize, srcCoord, dstCoord))

	//Part two
	stepsC := make(chan int)

	for y := 0; y < sideSize; y++ {
		for x := 0; x < sideSize; x++ {
			go func(x, y int) {
				if maze[y*sideSize+x] == -1 {
					stepsC <- -1
				} else {
					stepsC <- LeeAlgorithm(
						GenMaze(sideSize, input),
						sideSize,
						Coord{1, 1},
						Coord{x, y},
					)
				}
			}(x, y)
		}
	}

	locations := 0
	for i := 0; i < sideSize*sideSize; i++ {
		steps := <-stepsC
		if steps != -1 && steps <= 50 {
			locations++
		}
	}

	fmt.Println(locations)
}

func LeeAlgorithm(maze []int, sideSize int, src, dst Coord) int {
	i := 0
	dstCoord := dst.OneD(sideSize)
	maze[src.OneD(sideSize)] = -2

	uncheckedCoords := make([]Coord, 1)
	uncheckedCoords[0] = src

	check := func(p Coord, l *[]Coord, i int) {
		if p.x < sideSize && p.x >= 0 && p.y < sideSize && p.y >= 0 && maze[p.OneD(sideSize)] == 0 {
			maze[p.OneD(sideSize)] = i
			*l = append(*l, p)
		}
	}

	for maze[dstCoord] == 0 {
		i++

		l := make([]Coord, 0, len(uncheckedCoords)*4)
		for _, c := range uncheckedCoords {
			check(Coord{c.x - 1, c.y}, &l, i)
			check(Coord{c.x + 1, c.y}, &l, i)
			check(Coord{c.x, c.y - 1}, &l, i)
			check(Coord{c.x, c.y + 1}, &l, i)
		}

		if len(l) == 0 {
			return -1
		}
		uncheckedCoords = l
	}

	return i
}

func GenMaze(sideSize, magicNumber int) []int {
	maze := make([]int, sideSize*sideSize)
	for y := 0; y < sideSize; y++ {
		for x := 0; x < sideSize; x++ {
			p := x*x + 3*x + 2*x*y + y + y*y
			p += magicNumber
			bitsSet := CountSetBits(p)
			maze[y*sideSize+x] = (bitsSet % 2) * -1
		}
	}
	return maze
}

func CountSetBits(x int) int {
	bitsSet := 0
	for ; x != 0; x &= x - 1 {
		bitsSet++
	}
	return bitsSet
}
