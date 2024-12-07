package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func accumulate(target uint64, numbers []uint64, soFar uint64) bool {
	if soFar > target {
		return false
	}
	if soFar == target {
		// the solution (possibly erroneously) requires only parts
		// of the numbers, which makes close to no sense at all
		// the given example always uses the entirety of the numbers
		// while the accepted solution lets you stop before you
		// run out of numbers to combine with.
		return true //len(numbers) == 0
	}
	if len(numbers) == 0 {
		return false
	}
	next := numbers[0]
	if accumulate(target, numbers[1:], soFar+next) {
		return true
	}
	if accumulate(target, numbers[1:], soFar*next) {
		return true
	}
	return false
}

func accumulate2(target uint64, numbers []uint64, soFar uint64) bool {
	if soFar > target {
		return false
	}
	if soFar == target {
		return true //len(numbers) == 0
	}
	if len(numbers) == 0 {
		return false
	}
	next := numbers[0]
	if accumulate2(target, numbers[1:], soFar+next) {
		return true
	}
	if accumulate2(target, numbers[1:], soFar*next) {
		return true
	}
	p, e := strconv.ParseInt(
		fmt.Sprintf("%v%v", soFar, next),
		10, 64)
	if e != nil {
		return false
	}
	concat := uint64(p)
	if accumulate2(target, numbers[1:], concat) {
		return true
	}
	return false
}

type task struct {
	target  uint64
	numbers []uint64
}

func parseLine(line string) task {
	targetAndNumbers := strings.Split(line, ": ")
	target, err := strconv.ParseInt(targetAndNumbers[0], 10, 64)
	if err != nil {
		panic(err)
	}
	split := strings.Split(targetAndNumbers[1], " ")
	numbers := make([]uint64, len(split))
	for i := 0; i < len(split); i++ {
		num, err := strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}
		numbers[i] = uint64(num)
	}
	return task{
		target:  uint64(target),
		numbers: numbers,
	}
}

func parseInput() func(func(task) bool) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	return func(yield func(task) bool) {
		defer func() {
			_ = f.Close()
		}()
		for s.Scan() {
			if !yield(parseLine(s.Text())) {
				return
			}
		}
	}
}

func solve1() uint64 {
	var total uint64

	for res := range parseInput() {
		if accumulate(res.target, res.numbers[1:], res.numbers[0]) {
			total += res.target
		}
	}
	return total
}

func solve2() uint64 {
	var total uint64
	for res := range parseInput() {
		if accumulate2(res.target, res.numbers[1:], res.numbers[0]) {
			total += res.target
		}
	}
	return total
}

func main() {
	println(solve1())
	println(solve2())
}
