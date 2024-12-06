package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var direction = map[byte][2]int{
	'<': {0, -1},
	'>': {0, 1},
	'^': {-1, 0},
	'v': {1, 0},
}
var changes = map[byte]byte{
	'<': '^',
	'^': '>',
	'>': 'v',
	'v': '<',
}

type coord struct {
	x, y int
}

func doMoves(i, j, li, lj int, input [][]byte) (res map[coord][]coord, casCycle bool) {
	currentDir := input[i][j]
	res = make(map[coord][]coord)
	movesTo := direction[currentDir]
	nextI := i + movesTo[0]
	nextJ := j + movesTo[1]
	for {
		if !(nextI < li &&
			nextI > -1 &&
			nextJ < lj &&
			nextJ > -1) {
			// out of bounds
			break
		}
		c := coord{i, j}
		currentVisited := res[c]
		nextC := coord{nextI, nextJ}
		if slices.Contains(currentVisited, nextC) {
			return res, true
		}
		res[c] = append(currentVisited, nextC)

		if input[nextI][nextJ] == '#' {
			// rotate
			nextI = i
			nextJ = j
			currentDir = changes[currentDir]
			movesTo = direction[currentDir]
		} else {
			i, j = nextI, nextJ
		}

		nextI = i + movesTo[0]
		nextJ = j + movesTo[1]
	}
	c := coord{i, j}
	currentVisited := res[c]
	res[c] = append(currentVisited, coord{nextI, nextJ})
	return res, false
}

func countX(road map[coord][]coord) int {
	count := 0

	for range road {
		count++
	}
	return count
}

func findStartPos(input [][]byte) (start coord, li, lj int) {
	li = len(input)
	lj = len(input[0])

	// find guard pos
	for i := 0; i < li; i++ {
		for j := 0; j < lj; j++ {
			if _, ok := changes[input[i][j]]; ok {
				return coord{i, j}, li, lj
			}
		}
	}
	return coord{-1, -1}, li, lj
}

func solve1(start coord, li, lj int, input [][]byte) int {
	moves, _ := doMoves(start.x, start.y, li, lj, input)
	return countX(moves)
}

func parseInput() [][]byte {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(file)
	result := make([][]byte, 0)
	for s.Scan() {
		result = append(result, []byte(s.Text()))
	}
	return result
}

func solve2(start coord, li, lj int, input [][]byte) int {
	moves, _ := doMoves(start.x, start.y, li, lj, input)
	cyclesCount := 0
	for current := range moves {
		if current != start {
			input[current.x][current.y] = '#'
			_, hasCycles := doMoves(start.x, start.y, li, lj, input)
			if hasCycles {
				cyclesCount++
			}
			input[current.x][current.y] = '.'
		}
	}
	return cyclesCount
}

func main() {
	input := parseInput()
	start, li, lj := findStartPos(input)
	fmt.Println(solve1(start, li, lj, input))
	fmt.Println(solve2(start, li, lj, input))
}
