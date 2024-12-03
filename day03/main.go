package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lpappalettera/advent-of-code-2024/util"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(filename string) int {
	input := util.Read(filename)
	result := 0

	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	parts := r.FindAllString(input, -1)

	for _, it := range parts {
		c := strings.Index(it, ",")
		i1, _ := strconv.Atoi(it[4:c])
		i2, _ := strconv.Atoi(it[c+1 : len(it)-1])
		result += i1 * i2
	}

	return result
}

func part2(filename string) int {
	input := util.Read(filename)
	result := 0

	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	parts := r.FindAllString(input, -1)
	partsI := findIndexes(`mul\(([0-9]+),([0-9]+)\)`, input)

	dos := findIndexes(`do\(\)`, input)
	donts := findIndexes(`don\'t\(\)`, input)

	for i, pos := range partsI {
		part := parts[i]
		lastDo := lastCommand(dos, pos)
		lastDont := lastCommand(donts, pos)

		if lastDont == -1 || lastDo > lastDont {
			c := strings.Index(part, ",")
			i1, _ := strconv.Atoi(part[4:c])
			i2, _ := strconv.Atoi(part[c+1 : len(part)-1])

			result += i1 * i2
		}

	}

	return result
}

func lastCommand(items []int, max int) int {
	result := -1
	for _, it := range items {
		if it < max {
			result = it
		} else if it > max {
			break
		}
	}
	return result
}

func findIndexes(regex string, input string) []int {
	r, _ := regexp.Compile(regex)
	indexInfo := r.FindAllStringSubmatchIndex(input, -1)
	result := make([]int, len(indexInfo))

	for i, it := range indexInfo {
		result[i] = it[0]
	}

	return result
}
