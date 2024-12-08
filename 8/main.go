package main

import (
	"bufio"
	"fmt"
	"os"
)

type c2d struct {
	x, y int
}

type board struct {
	antennaePositions map[byte][]c2d
	antiNodes         map[c2d]bool
	dims              c2d
}

func newBoard(theirs bool) board {
	fileName := "input.txt"
	if theirs {
		fileName = "theirs.txt"
	}
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)
	rows := 0
	cols := 0
	b := board{
		antennaePositions: make(map[byte][]c2d),
		antiNodes:         make(map[c2d]bool),
	}
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if cols == 0 {
			cols = len(line)
		}
		for pos, v := range line {
			if v != '.' {
				current := b.antennaePositions[v]
				b.antennaePositions[v] = append(current, c2d{
					x: rows,
					y: pos,
				})
			}
		}
		rows++
	}
	b.dims.x = rows
	b.dims.y = cols
	return b
}

func (a *c2d) next(b c2d) c2d {
	return c2d{
		a.x - (b.x - a.x),
		a.y - (b.y - a.y),
	}
}

func (a *c2d) nextN(b c2d, n int) c2d {
	return c2d{
		a.x - n*(b.x-a.x),
		a.y - n*(b.y-a.y),
	}
}

func (a *c2d) isValid(dims c2d) bool {
	return a.x >= 0 && a.y >= 0 && a.x < dims.x && a.y < dims.y
}

func (b *board) findAntiNodes() {
	for _, coords := range b.antennaePositions {
		lc := len(coords)
		for i := 0; i < lc-1; i++ {
			for j := i + 1; j < lc; j++ {
				ij := coords[i].next(coords[j])
				if ij.isValid(b.dims) {
					b.antiNodes[ij] = true
				}
				ji := coords[j].next(coords[i])
				if ji.isValid(b.dims) {
					b.antiNodes[ji] = true
				}
			}
		}
	}
}

func solve2(b *board) int {
	antiNodes := make(map[c2d]bool)
	for _, coords := range b.antennaePositions {
		lc := len(coords)
		for i := 0; i < lc-1; i++ {
			for j := i + 1; j < lc; j++ {
				for n := 0; ; n++ {
					ij := coords[i].nextN(coords[j], n)
					added := false
					if ij.isValid(b.dims) {
						antiNodes[ij] = true
						added = true
					}
					ji := coords[j].nextN(coords[i], n)
					if ji.isValid(b.dims) {
						antiNodes[ji] = true
						added = true
					}
					if !added {
						break
					}
				}
			}
		}
	}
	return len(antiNodes)
}

func solve1(b *board) int {
	return len(b.antiNodes)
}

func main() {
	b := newBoard(false)
	b.findAntiNodes()
	fmt.Println(solve1(&b))
	fmt.Println(solve2(&b))
}
