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
	result := 0

	for _, levels := range parse(filename) {
		if isSafe(levels) {
			result++
		}
	}

	return result
}

func part2(filename string) int {
	result := 0

	for _, levels := range parse(filename) {
		safe := isSafe(levels)

		if !safe {
			for i := range levels {
				if isSafe(util.RemoveIndex(levels, i)) {
					safe = true
					break
				}
			}
		}

		if safe {
			result++
		}
	}

	return result
}

func isSafe(levels []int) bool {
	safe := true
	inc := levels[0] < levels[1]

	for i := 1; i < len(levels); i++ {
		d := 0
		if inc {
			d = levels[i] - levels[i-1]
		} else {
			d = levels[i-1] - levels[i]
		}

		if d < 1 || d > 3 {
			safe = false
			break
		}
	}

	return safe
}

func parse(filename string) [][]int {
	lines := util.ReadLines(filename)
	var result [][]int

	for _, line := range lines {
		parts := strings.Split(line, " ")
		var levels []int

		for _, part := range parts {
			level, err := strconv.Atoi(part)
			util.HandleError(err)

			levels = append(levels, level)
		}

		result = append(result, levels)
	}

	return result
}
