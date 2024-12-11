package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stone int64

func (s stone) split() (first, second stone, success bool) {
	digits := make([]stone, 0)
	first = s
	for s > 0 {
		digits = append(digits, s%10)
		s /= 10
	}
	l := len(digits)
	if l%2 == 1 {
		return first, 0, false
	}
	midpoint := l / 2

	// 0 1 2 3 --> len = 4
	// first = 01 (start from midpoint - 1, all the way to 0)
	first = digits[midpoint-1]
	for i := midpoint - 2; i >= 0; i-- {
		first *= 10
		first += digits[i]
	}
	// second = 23 (start from len - 1, all the way to midpoint)
	second = digits[l-1]
	for i := l - 2; i >= midpoint; i-- {
		second *= 10
		second += digits[i]
	}
	return second, first, true
}

func (s stone) change() (first, second stone, split bool) {
	if s == 0 {
		return 1, 0, false
	}
	first, second, split = s.split()
	if split {
		return first, second, true
	}
	return first * 2024, 0, false
}

type stoneAtStep struct {
	step int
	s    stone
}

var cache = make(map[stoneAtStep]stone)

func (s stone) stonesAfterBlinks(count int) stone {
	recurse := func(s stone, count int) stone {
		callIndex := stoneAtStep{step: count, s: s}
		result, ok := cache[callIndex]
		if ok {
			return result
		}
		if count == 0 {
			cache[stoneAtStep{step: count, s: s}] = 1
			return 1
		}
		if s == 0 {
			result = stone(1).stonesAfterBlinks(count - 1)
			cache[callIndex] = result
			return result
		}
		first, second, split := s.change()
		if split {
			f := stoneAtStep{count - 1, first}
			s1, ok := cache[f]
			if !ok {
				s1 = first.stonesAfterBlinks(count - 1)
				cache[f] = s1
			}
			ss := stoneAtStep{count - 1, second}
			s2, ok := cache[ss]
			if !ok {
				s2 = second.stonesAfterBlinks(count - 1)
				cache[ss] = s2
			}
			cache[callIndex] = s1 + s2
			return s1 + s2
		}
		result = first.stonesAfterBlinks(count - 1)
		cache[callIndex] = result
		return result
	}
	return recurse(s, count)
}

func parseInput(theirs bool) []stone {
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
	scanner.Scan()
	text := scanner.Text()
	numbers := strings.Split(text, " ")
	result := make([]stone, len(numbers))
	for i := 0; i < len(numbers); i++ {
		num, err := strconv.ParseInt(numbers[i], 10, 64)
		if err != nil {
			panic(err)
		}
		result[i] = stone(num)
	}
	return result
}

func solve1(input []stone) stone {
	var result stone = 0
	for _, s := range input {
		result += s.stonesAfterBlinks(25)
	}
	return result
}

func solve2(input []stone) stone {
	var result stone = 0
	for _, s := range input {
		result += s.stonesAfterBlinks(75)
	}
	return result
}

func main() {
	input := parseInput(false)
	fmt.Println(solve1(input))
	fmt.Println(solve2(input))
}
