package main

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x, y int64
}

type machine struct {
	a, b  coord
	prize coord
}

func (m machine) cost(offset int64) int64 {
	var prizeX = m.prize.x + offset
	var prizeY = m.prize.y + offset
	det := m.a.x*m.b.y - m.a.y*m.b.x
	a := (prizeX*m.b.y - prizeY*m.b.x) / det
	b := (m.a.x*prizeY - m.a.y*prizeX) / det
	if m.a.x*a+m.b.x*b == prizeX && m.a.y*a+m.b.y*b == prizeY {
		return a*3 + b
	}
	return 0
}

func parseMachine(s *bufio.Scanner) *machine {
	m := machine{}
	var x, y int64

	if !s.Scan() {
		return nil
	}
	// read button A
	line := s.Text()
	if len(line) == 0 {
		return nil
	}
	n, err := fmt.Sscanf(line, "Button A: X+%v, Y+%v", &x, &y)
	if err != nil || n != 2 {
		panic("Cannot parse button a")
	}
	m.a = coord{x, y}

	// read button B
	if !s.Scan() {
		panic("Cannot scan button b")
	}
	line = s.Text()
	n, err = fmt.Sscanf(line, "Button B: X+%v, Y+%v", &x, &y)
	if err != nil || n != 2 {
		panic("Cannot parse button b")
	}
	m.b = coord{x, y}

	// read prize
	if !s.Scan() {
		panic("Cannot scan prize")
	}
	line = s.Text()
	n, err = fmt.Sscanf(line, "Prize: X=%v, Y=%v", &x, &y)
	if err != nil || n != 2 {
		panic("Cannot parse button prize")
	}
	m.prize = coord{x, y}
	return &m
}

func parseInput(theirs bool) func(yield func(m machine) bool) {
	fileName := "input.txt"
	if theirs {
		fileName = "theirs.txt"
	}
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	return func(yield func(m machine) bool) {
		defer func() {
			_ = f.Close()
		}()

		for {
			m := parseMachine(scanner)
			if m == nil {
				break
			}
			// move one line down
			scanner.Scan()
			if !yield(*m) {
				break
			}
		}
	}
}

func solve1(machineStream func(yield func(m machine) bool)) int64 {
	var result int64 = 0
	for m := range machineStream {
		result += m.cost(0)
	}
	return result
}

func solve2(machineStream func(yield func(m machine) bool)) (result int64) {
	result = 0
	for m := range machineStream {
		result += m.cost(10000000000000)
	}
	return result
}

func main() {
	input := parseInput(false)
	fmt.Println(solve1(input))
	input = parseInput(false)
	fmt.Println(solve2(input))
}
