package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseRules(rules []string) map[int]map[int]bool {
	result := make(map[int]map[int]bool)
	for _, rule := range rules {
		split := strings.Split(rule, "|")
		first, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		if result[first] == nil {
			result[first] = make(map[int]bool)
		}
		result[first][second] = true
	}
	return result
}

func isInCorrectOrder(printout []int, rules map[int]map[int]bool) bool {
	for index, element := range printout {
		rulesForElement, ok := rules[element]
		if !ok {
			continue
		}
		for j := 0; j < index; j++ {
			if rulesForElement[printout[j]] {
				return false
			}
		}
	}
	return true
}

func solve1(printouts [][]int, rules map[int]map[int]bool) int {
	result := 0
	for _, printout := range printouts {
		if isInCorrectOrder(printout, rules) {
			midElement := printout[len(printout)/2]
			result += midElement
		}
	}
	return result
}

type customSort struct {
	data  []int
	rules map[int]map[int]bool
}

func (s *customSort) Len() int {
	return len(s.data)
}
func (s *customSort) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}
func (s *customSort) Less(i, j int) bool {
	return s.rules[s.data[i]][s.data[j]]
}

func solve2(printouts [][]int, rules map[int]map[int]bool) int {
	result := 0
	bubu := customSort{data: []int{}, rules: rules}
	for _, printout := range printouts {
		if isInCorrectOrder(printout, rules) {
			continue
		}
		bubu.data = printout
		sort.Sort(&bubu)
		result += bubu.data[bubu.Len()/2]
	}
	return result
}

func parseInput() (printouts [][]int, rules map[int]map[int]bool) {
	file, err := os.Open("input.txt")
	defer func() {
		_ = file.Close()
	}()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	printouts = make([][]int, 0)
	notParsedRules := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break // now printouts
		}
		notParsedRules = append(notParsedRules, line)
	}
	for scanner.Scan() {
		line := scanner.Text()
		stringPages := strings.Split(line, ",")
		pages := make([]int, len(stringPages))
		for i, page := range stringPages {
			pages[i], err = strconv.Atoi(page)
			if err != nil {
				panic(err)
			}
		}
		printouts = append(printouts, pages)
	}
	return printouts, parseRules(notParsedRules)
}

func main() {
	inputs, rules := parseInput()
	fmt.Println(solve1(inputs, rules))
	fmt.Println(solve2(inputs, rules))
}
