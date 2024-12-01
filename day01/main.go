package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/lpappalettera/advent-of-code-2024/util"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(filename string) int {
	a, b := parse(filename)
	result := 0

	slices.Sort(a)
	slices.Sort(b)

	for i, i1 := range a {
		i2 := b[i]
		if i2 > i1 {
			result += i2 - i1
		} else {
			result += i1 - i2
		}
	}

	return result
}

func part2(filename string) int {
	a, b := parse(filename)
	result := 0

	for _, i1 := range a {
		count := 0

		for _, i2 := range b {
			if i2 == i1 {
				count++
			} 
		}

		result += count * i1
	}

	return result
}

func parse(filename string) ([]int, []int) {
	lines := util.ReadLines(filename)
	var a []int
	var b []int

	for _, line := range lines {
		parts := strings.Split(line, "   ")

		i1, err := strconv.Atoi(parts[0])
		util.HandleError(err)
		i2, err := strconv.Atoi(parts[1])
		util.HandleError(err)

		a = append(a, i1)
		b = append(b, i2)
	}

	return a, b
}
