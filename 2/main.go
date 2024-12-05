package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func isSafe(input []int) bool {
	l := len(input)
	if input[0] > input[1] {
		// rising
		for i := 1; i < l; i++ {
			diff := input[i-1] - input[i]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	} else if input[0] < input[1] {
		for i := 1; i < l; i++ {
			diff := input[i] - input[i-1]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func isSafe2(input []int) bool {
	l := len(input)
	failedOnce := false
	if input[0] > input[1] {
		// rising
		for i := 1; i < l; i++ {
			diff := input[i-1] - input[i]
			if diff < 1 || diff > 3 {
				if failedOnce {
					return false
				} else {
					failedOnce = true
				}
			}
		}
	} else if input[0] < input[1] {
		for i := 1; i < l; i++ {
			diff := input[i] - input[i-1]
			if diff < 1 || diff > 3 {
				if failedOnce {
					return false
				} else {
					failedOnce = true
				}
			}
		}
	} else {
		return false
	}
	return true
}

func readInput1() [][]int {
	f, err := os.Open("my.txt")
	res := make([][]int, 0)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		numbers := strings.Split(line, " ")
		input := make([]int, len(numbers))

		for i, v := range numbers {
			input[i], err = strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
		}
		res = append(res, input)
	}
	return res
}

func solve1() int {
	input := readInput1()
	total := 0
	for _, line := range input {
		if isSafe(line) {
			total++
		}
	}
	return total
}

func solve2() int {
	input := readInput1()
	total := 0
	for _, line := range input {
		if isSafe2(line) {
			total++
		}
	}
	return total
}

func main() {
	println(solve1())
	println(solve2())
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
