package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int
}

type input struct {
	maxX, maxY int
	data       [][]byte
}

func (i *input) getStartingPositions() func(yield func(x, y int) bool) {
	return func(yield func(x, y int) bool) {
		for rowId, row := range i.data {
			for colId, cell := range row {
				if cell == 0 {
					if !yield(rowId, colId) {
						return
					}
				}
			}
		}
	}
}

func (i *input) getTrailheadsCount(
	x, y int,
	visited map[coord]bool,
) int {
	cell := i.data[x][y]
	if cell == 9 {
		c := coord{
			x: x,
			y: y,
		}
		if visited[c] {
			return 0
		} else {
			visited[c] = true
			return 1
		}
	}
	next := cell + 1
	trailheadsCount := 0
	// try moving up
	if x > 0 && i.data[x-1][y] == next {
		trailheadsCount += i.getTrailheadsCount(x-1, y, visited)
	}
	// try moving down
	if x < i.maxX-1 && i.data[x+1][y] == next {
		trailheadsCount += i.getTrailheadsCount(x+1, y, visited)
	}
	// try moving left
	if y > 0 && i.data[x][y-1] == next {
		trailheadsCount += i.getTrailheadsCount(x, y-1, visited)
	}
	if y < i.maxY-1 && i.data[x][y+1] == next {
		trailheadsCount += i.getTrailheadsCount(x, y+1, visited)
	}
	return trailheadsCount
}

func (i *input) getDistinctHikingTrails(x, y int) int {
	cell := i.data[x][y]
	if cell == 9 {
		return 1
	}
	next := cell + 1
	trailheadsCount := 0
	// try moving up
	if x > 0 && i.data[x-1][y] == next {
		trailheadsCount += i.getDistinctHikingTrails(x-1, y)
	}
	// try moving down
	if x < i.maxX-1 && i.data[x+1][y] == next {
		trailheadsCount += i.getDistinctHikingTrails(x+1, y)
	}
	// try moving left
	if y > 0 && i.data[x][y-1] == next {
		trailheadsCount += i.getDistinctHikingTrails(x, y-1)
	}
	if y < i.maxY-1 && i.data[x][y+1] == next {
		trailheadsCount += i.getDistinctHikingTrails(x, y+1)
	}
	return trailheadsCount
}

func solve1(input input) int {
	total := 0
	for x, y := range input.getStartingPositions() {
		total += input.getTrailheadsCount(x, y, make(map[coord]bool))
	}
	return total
}

func solve2(input input) int {
	total := 0
	for x, y := range input.getStartingPositions() {
		total += input.getDistinctHikingTrails(x, y)
	}
	return total
}

func parseInput(theirs bool) input {
	fileName := "input.txt"
	if theirs {
		fileName = "theirs.txt"
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	result := input{
		data: make([][]byte, 0),
	}
	for scanner.Scan() {
		line := []byte(scanner.Text())
		for i, v := range line {
			line[i] = v - '0'
		}
		result.data = append(result.data, line)
		lineCount++
	}
	result.maxX = lineCount
	result.maxY = len(result.data[0])
	return result
}

func main() {
	input := parseInput(false)
	fmt.Println(solve1(input))
	fmt.Println(solve2(input))
}
