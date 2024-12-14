package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
)

type coord struct {
	x, y int64
}

func (c coord) project(dims coord) coord {
	x := ((c.x % dims.x) + dims.x) % dims.x
	y := ((c.y % dims.y) + dims.y) % dims.y
	return coord{
		x: x,
		y: y,
	}
}

type robot struct {
	initialPosition coord
	velocity        coord
}

func (r robot) step(n int64) coord {
	return coord{
		x: r.initialPosition.x + n*r.velocity.x,
		y: r.initialPosition.y + n*r.velocity.y,
	}
}

func parseInput(theirs bool) []robot {
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
	result := make([]robot, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, vx, vy int64
		num, err := fmt.Sscanf(line, "p=%v,%v v=%v,%v", &x, &y, &vx, &vy)
		if err != nil {
			panic(err)
		}
		if num < 4 {
			panic("not enough values scanned")
		}
		r := robot{
			initialPosition: coord{
				x: x,
				y: y,
			},
			velocity: coord{
				x: vx,
				y: vy,
			},
		}
		result = append(result, r)
	}
	return result
}

func solve1(robots []robot, dims coord) int64 {
	quadrants := [4]int64{}
	x2 := dims.x / 2
	y2 := dims.y / 2
	for _, r := range robots {
		c := r.step(100).project(dims)
		if c.x < x2 && c.y < y2 {
			quadrants[0]++
		} else if c.x < x2 && c.y > y2 {
			quadrants[1]++
		} else if c.x > x2 && c.y < y2 {
			quadrants[2]++
		} else if c.x > x2 && c.y > y2 {
			quadrants[3]++
		}
	}
	var result int64 = 1
	for _, v := range quadrants {
		result *= v
	}
	return result
}

func solve2(robots []robot, dims coord) {
	// french automation
	// half the solution has to be done manually :)
	for i := 0; i < 10000; i++ {
		img := image.NewGray(image.Rect(0, 0, int(dims.x), int(dims.y)))
		for _, r := range robots {
			c := r.step(int64(i)).project(dims)
			x := int(c.x)
			y := int(c.y)
			img.Set(x, y, image.White)
		}
		f, err := os.OpenFile(fmt.Sprintf("img/frame%v.png", i), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		err = png.Encode(f, img)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	theirs := false
	dims := coord{
		x: 101,
		y: 103,
	}
	if theirs {
		dims = coord{
			x: 11,
			y: 7,
		}
	}
	input := parseInput(theirs)
	fmt.Println(solve1(input, dims))
	solve2(input, dims)
}
