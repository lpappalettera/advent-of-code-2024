package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lpappalettera/advent-of-code-2024/util"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(filename string) int {
	input := parse(filename)
	result := 0

	for _, line := range input {
		if calc(line.x, line.values[0], line.values[1:], false) {
			result += line.x
		}
	}

	return result
}

func part2(filename string) int {
	input := parse(filename)
	result := 0

	for _, line := range input {
		if calc(line.x, line.values[0], line.values[1:], true) {
			result += line.x
		}
	}

	return result
}

func calc(x int, val int, values []int, c bool) bool {
	if len(values) > 1 {
		return calc(x, val*values[0], values[1:], c) || calc(x, val+values[0], values[1:], c) || (c && calc(x, concat(val, values[0]), values[1:], c))
	} else if val*values[0] == x {
		return true
	} else if val+values[0] == x {
		return true
	} else if c && concat(val, values[0]) == x {
		return true
	}

	return false
}

func concat(a int, b int) int {
	return util.StrToInt(strconv.Itoa(a) + strconv.Itoa(b))
}

type input struct {
	x      int
	values []int
}

func parse(filename string) []input {
	lines := util.ReadLines(filename)
	var result []input

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result = append(result, input{
			x:      util.StrToInt(parts[0]),
			values: util.Map(strings.Split(parts[1], " "), util.StrToInt),
		})
	}

	return result
}
