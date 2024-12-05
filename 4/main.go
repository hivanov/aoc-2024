package main

import (
	"bufio"
	"fmt"
	"os"
)

type state struct {
	mode  int
	found int
}

func (s *state) next(c byte) bool {
	if s.mode == 0 {
		if c == 'M' {
			s.mode = 1
			return true
		}
	} else if s.mode == 1 {
		if c == 'A' {
			s.mode = 2
			return true
		}
	} else if s.mode == 2 {
		if c == 'S' {
			s.found = 1
		}
	}
	return false
}

func search(xRow, xCol, lenRow, lenCol int, input [][]byte) int {
	total := 0
	// right
	if xCol < lenCol-3 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow][xCol+i]) {
				break
			}
		}
		total += s.found
	}

	// left
	if xCol > 2 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow][xCol-i]) {
				break
			}
		}
		total += s.found
	}

	// down
	if xRow < lenRow-3 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow+i][xCol]) {
				break
			}
		}
		total += s.found
	}

	// up
	if xRow > 2 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow-i][xCol]) {
				break
			}
		}
		total += s.found
	}

	// down-right
	if xRow < lenRow-3 && xCol < lenCol-3 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow+i][xCol+i]) {
				break
			}
		}
		total += s.found
	}

	// down-left
	if xRow < lenRow-3 && xCol > 2 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow+i][xCol-i]) {
				break
			}
		}
		total += s.found
	}
	// up-left
	if xRow > 2 && xCol > 2 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow-i][xCol-i]) {
				break
			}
		}
		total += s.found
	}

	// up-right
	if xRow > 2 && xCol < lenRow-3 {
		s := new(state)
		for i := 1; i < 4; i++ {
			if !s.next(input[xRow-i][xCol+i]) {
				break
			}
		}
		total += s.found
	}
	return total
}

func solve1(input [][]byte) int {
	total := 0
	li := len(input)
	lj := len(input[0])
	for i := 0; i < li; i++ {
		for j := 0; j < lj; j++ {
			if input[i][j] == 'X' {
				s := search(i, j, li, lj, input)
				total += s
			}
		}
	}
	return total
}

var sam = "SAM"
var mas = "MAS"

func starCheck(a, b string) bool {
	return (a == sam || a == mas) && (b == mas || b == sam)
}

func solve2(input [][]byte) int {
	total := 0
	li := len(input)
	lj := len(input[0])
	for i := 1; i < li-1; i++ {
		for j := 1; j < lj-1; j++ {
			a := [3]byte{input[i-1][j-1], input[i][j], input[i+1][j+1]}
			b := [3]byte{input[i-1][j+1], input[i][j], input[i+1][j-1]}
			if starCheck(string(a[:]), string(b[:])) {
				total++
			}
			/*
				a = [3]byte{input[i-1][j], input[i][j], input[i+1][j]}
				b = [3]byte{input[i][j-1], input[i][j], input[i][j+1]}
				if starCheck(string(a[:]), string(b[:])) {
					total++
				}
			*/
		}
	}
	return total
}

func readInput() ([][]byte, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	s := bufio.NewScanner(f)
	result := make([][]byte, 0)
	for s.Scan() {
		result = append(result, []byte(s.Text()))
	}
	return result, nil
}

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	fmt.Println(solve1(input))
	fmt.Println(solve2(input))
}
