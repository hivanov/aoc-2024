package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solve1(input []byte) (uint64, error) {
	r, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	if err != nil {
		return 0, err
	}
	var result uint64 = 0
	for _, v := range r.FindAllSubmatch(input, -1) {
		a, err := strconv.Atoi(string(v[1]))
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(string(v[2]))
		if err != nil {
			return 0, err
		}
		result += uint64(a * b)
	}
	return result, nil
}

func solve2(input []byte) uint64 {
	mode := 0
	isEnabled := 1
	a := 0
	b := 0
	var result uint64 = 0
	for _, v := range input {
		if mode == 0 { // no word
			if v == 'm' {
				mode = 1 // mul?
			} else if v == 'd' { // do?
				mode = 10
			}
		} else if mode == 1 { // mu?
			if v == 'u' {
				mode = 2
			} else {
				mode = 0
			}
		} else if mode == 2 {
			if v == 'l' {
				mode = 3
			} else {
				mode = 0
			}
		} else if mode == 3 {
			if v == '(' {
				mode = 4
			} else {
				mode = 0
			}
		} else if mode == 4 {
			if v >= '0' && v <= '9' {
				a = a*10 + int(v-'0')
			} else if v == ',' { // b now
				mode = 5
			} else {
				mode = 0
			}
		} else if mode == 5 {
			if v >= '0' && v <= '9' {
				b = b*10 + int(v-'0')
			} else if v == ')' {
				result += uint64(a * b * isEnabled)
				mode = 0
			} else {
				mode = 0
			}
		} else if mode == 10 {
			if v == 'o' {
				mode = 11
			} else {
				mode = 0
			}
		} else if mode == 11 {
			if v == '(' {
				mode = 12
			} else if v == 'n' {
				mode = 20
			} else {
				mode = 0
			}
		} else if mode == 12 {
			if v == ')' {
				isEnabled = 1
			}
			mode = 0
		} else if mode == 20 {
			if v == '\'' {
				mode = 21
			} else {
				mode = 0
			}
		} else if mode == 21 {
			if v == 't' {
				mode = 22
			} else {
				mode = 0
			}
		} else if mode == 22 {
			if v == '(' {
				mode = 23
			} else {
				mode = 0
			}
		} else if mode == 23 {
			if v == ')' {
				isEnabled = 0
			}
			mode = 0
		}
		if mode == 0 {
			a = 0
			b = 0
		}
	}
	return result
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	res, err := solve1(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(solve2(f))
}
