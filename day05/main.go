package main

import (
	"fmt"
	"slices"
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

	for _, row := range input.pages {
		valid := true

		for _, rule := range input.rules {
			a := slices.Index(row, rule[0])
			if a == -1 {
				continue
			}

			b := slices.Index(row, rule[1])
			if b != -1 && a > b {
				valid = false
				break
			}
		}

		if valid {
			result += row[len(row)/2]
		}

	}

	return result
}

func part2(filename string) int {
	input := parse(filename)
	result := 0

	for _, row := range input.pages {
		valid := true
		rdy := false

		for !rdy {
			updated := false
			for a, page := range row {
				for _, rule := range input.rules {
					if rule[0] != page {
						continue
					}

					b := slices.Index(row, rule[1])
					if b != -1 && a > b {
						valid = false
						updated = true
						row[a], row[b] = row[b], row[a]
					}
				}
			}

			if !updated {
				rdy = true
			}
		}

		if !valid {
			result += row[len(row)/2]
		}

	}

	return result
}

type input struct {
	rules [][]int
	pages [][]int
}

func parse(filename string) input {
	lines := util.ReadLines(filename)
	p := 0
	var rules [][]int
	var pages [][]int

	for _, line := range lines {
		if line == "" {
			p++
		} else if p == 0 {
			parts := util.Map(strings.Split(line, "|"), util.StrToInt)
			rules = append(rules, []int{parts[0], parts[1]})
		} else if p == 1 {
			items := util.Map(strings.Split(line, ","), util.StrToInt)
			pages = append(pages, items)
		}
	}

	return input{rules, pages}
}
