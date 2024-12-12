package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type square struct {
	visited bool
	cropId  byte
}

type coord struct {
	x, y int
}

func (c coord) up() coord {
	return coord{c.x - 1, c.y}
}

func (c coord) down() coord {
	return coord{c.x + 1, c.y}
}

func (c coord) left() coord {
	return coord{c.x, c.y - 1}
}

func (c coord) right() coord {
	return coord{c.x, c.y + 1}
}

func (c coord) inBounds(maxX, maxY int) bool {
	return c.x >= 0 && c.x < maxX && c.y >= 0 && c.y < maxY
}

type plot map[coord]bool

func (p plot) area() int {
	return len(p)
}

func (p plot) perimeter() int {
	result := 0
	for c := range p {
		outwardFacingSides := 4
		if _, ok := p[c.up()]; ok {
			outwardFacingSides--
		}
		if _, ok := p[c.down()]; ok {
			outwardFacingSides--
		}
		if _, ok := p[c.left()]; ok {
			outwardFacingSides--
		}
		if _, ok := p[c.right()]; ok {
			outwardFacingSides--
		}
		result += outwardFacingSides
	}
	return result
}

type coords []coord

func (v coords) Len() int { return len(v) }
func (v coords) Less(i, j int) bool {
	if v[i].x == v[j].x {
		return v[i].y < v[j].y
	} else if v[i].x < v[j].x {
		return true
	}
	return false
}
func (v coords) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (p plot) numberOfSides() int {
	leftFence := make(map[coord]int)
	rightFence := make(map[coord]int)
	downFence := make(map[coord]int)
	upFence := make(map[coord]int)
	fenceId := 1

	// "join previous fence, if possible" algorithm
	// usually, when trying to join an up/down fence, we search for "leftmost"
	// or "rightmost" fence already visited.
	// going from top-left to bottom-right, one row at a time (left-to-right),
	// first try to join the up/down fence originating from the plot on the
	// left. Then, the one of the plot on the right.
	//
	// Same thing happens for left/right fences, but the search is "top" to
	// "bottom"
	idOrNew := func(c, firstTry, secondTry coord, fenceSoFar map[coord]int) {
		if _, ok := p[firstTry]; ok {
			// the first try coordinate is in the plot
			if first, ok := fenceSoFar[firstTry]; ok {
				fenceSoFar[c] = first
				return
			}
		}
		if _, ok := p[secondTry]; ok {
			// the second try coordinate is in the plot
			if second, ok := fenceSoFar[secondTry]; ok {
				fenceSoFar[c] = second
				return
			}
		}
		fenceSoFar[c] = fenceId
		fenceId++
	}

	// we need to sort the coordinates of the area for the method below to work
	sorted := make(coords, len(p))
	i := 0
	for c := range p {
		sorted[i] = c
		i++
	}
	sort.Sort(sorted)

	for _, c := range sorted {
		left := c.left()
		right := c.right()
		up := c.up()
		down := c.down()
		if _, ok := p[up]; !ok {
			idOrNew(c, left, right, upFence)
		}
		if _, ok := p[down]; !ok {
			idOrNew(c, right, left, downFence)
		}
		if _, ok := p[left]; !ok {
			idOrNew(c, up, down, leftFence)
		}
		if _, ok := p[right]; !ok {
			idOrNew(c, up, down, rightFence)
		}
	}

	// count distinct fence ids
	fences := make(map[int]bool)
	for _, v := range leftFence {
		fences[v] = true
	}
	for _, v := range rightFence {
		fences[v] = true
	}
	for _, v := range upFence {
		fences[v] = true
	}
	for _, v := range downFence {
		fences[v] = true
	}
	return len(fences)

}

func (p plot) price() int {
	return p.perimeter() * p.area()
}

type plots map[int]plot

type garden struct {
	area       [][]square
	maxX, maxY int
}

func (g *garden) visit(c coord, p plot) {
	cell := g.area[c.x][c.y]
	if cell.visited {
		return
	}
	g.area[c.x][c.y].visited = true

	// add to current plot
	p[c] = true

	newC := c.up()
	if newC.inBounds(g.maxX, g.maxY) {
		newCell := g.area[newC.x][newC.y]
		if newCell.cropId == cell.cropId && !newCell.visited {
			g.visit(newC, p)
		}
	}
	newC = c.down()
	if newC.inBounds(g.maxX, g.maxY) {
		newCell := g.area[newC.x][newC.y]
		if newCell.cropId == cell.cropId && !newCell.visited {
			g.visit(newC, p)
		}
	}
	newC = c.left()
	if newC.inBounds(g.maxX, g.maxY) {
		newCell := g.area[newC.x][newC.y]
		if newCell.cropId == cell.cropId && !newCell.visited {
			g.visit(newC, p)
		}
	}
	newC = c.right()
	if newC.inBounds(g.maxX, g.maxY) {
		newCell := g.area[newC.x][newC.y]
		if newCell.cropId == cell.cropId && !newCell.visited {
			g.visit(newC, p)
		}
	}
}

func (g *garden) divideToPlots() plots {
	plotId := 0
	result := make(plots)
	for r, row := range g.area {
		for c := range row {
			if g.area[r][c].visited {
				continue
			}
			p := make(plot)
			g.visit(coord{r, c}, p)
			result[plotId] = p
			plotId++
		}
	}
	return result
}

func parseInput(theirs bool) garden {
	fileName := "input.txt"
	if theirs {
		fileName = "theirs.txt"
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	result := garden{
		area: make([][]square, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()
		result.maxY = len(line)
		row := make([]square, result.maxY)
		for i := range row {
			row[i].cropId = line[i]
		}
		result.area = append(result.area, row)
		result.maxX++
	}
	return result
}

func solve1(p plots) int {
	result := 0
	for _, pp := range p {
		result += pp.price()
	}
	return result
}

func solve2(p plots) int {
	result := 0
	for _, pp := range p {
		result += pp.area() * pp.numberOfSides()
	}
	return result
}

func main() {
	garden := parseInput(false)
	plots := garden.divideToPlots()
	fmt.Println(solve1(plots))
	fmt.Println(solve2(plots))
}
