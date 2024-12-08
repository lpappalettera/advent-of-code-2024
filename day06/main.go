package main

import (
	"fmt"

	"github.com/lpappalettera/advent-of-code-2024/util"
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(filename string) int {
	input, y, x := parse(filename)
	finished := false
	var result []coor

	for !finished {
		direction := input[y][x]

		unique := true
		for _, step := range result {
			if step.x == x && step.y == y {
				unique = false
			}
		}

		if unique {
			result = append(result, coor{y, x, direction})
		}

		nextStep(input, &y, &x, direction, &finished)
	}

	return len(result)
}

func part2(filename string) int {
	input, y, x := parse(filename)

	result := 0

	for copyY, line := range input {
		for copyX := range line {
			tmp := inputCopy(input)

			if tmp[copyY][copyX] == '.' {
				tmp[copyY][copyX] = '#'
				if hasLoop(tmp, y, x) {
					result++
				}
			}
		}
	}

	return result
}

func hasLoop(m [][]rune, y int, x int) bool {
	finished := false
	var steps []coor
	looped := false

	for !finished {
		direction := m[y][x]

		unique := true
		for _, step := range steps {
			if step.x == x && step.y == y && step.direction == direction {
				unique = false
				looped = true
				finished = true
			}
		}

		if unique {
			steps = append(steps, coor{y, x, direction})
		}

		nextStep(m, &y, &x, direction, &finished)
	}

	return looped
}

type coor struct {
	y         int
	x         int
	direction rune
}

func parse(filename string) ([][]rune, int, int) {
	lines := util.ReadLines(filename)
	var result [][]rune
	var startY, startX int

	for y, line := range lines {
		var row []rune
		for x, char := range line {
			if char == '<' || char == '^' || char == '>' || char == 'v' {
				startY, startX = y, x
			}

			row = append(row, char)
		}
		result = append(result, row)
	}

	return result, startY, startX
}

func nextStep(input [][]rune, y *int, x *int, direction rune, finished *bool) {
	switch direction {
	case '^':
		if !inBouds(input, *y-1, *x) {
			*finished = true
		} else if input[*y-1][*x] == '#' {
			input[*y][*x] = '>'
		} else {
			input[*y-1][*x] = '^'
			*y--
		}
	case '>':
		if !inBouds(input, *y, *x+1) {
			*finished = true
		} else if input[*y][*x+1] == '#' {
			input[*y][*x] = 'v'
		} else {
			input[*y][*x+1] = '>'
			*x++
		}
	case 'v':
		if !inBouds(input, *y+1, *x) {
			*finished = true
		} else if input[*y+1][*x] == '#' {
			input[*y][*x] = '<'
		} else {
			input[*y+1][*x] = 'v'
			*y++
		}
	case '<':
		if !inBouds(input, *y, *x-1) {
			*finished = true
		} else if input[*y][*x-1] == '#' {
			input[*y][*x] = '^'
		} else {
			input[*y][*x-1] = '<'
			*x--
		}
	}
}

func inBouds(input [][]rune, y int, x int) bool {
	return y >= 0 && x >= 0 && y < len(input) && x < len(input[0])
}

func inputCopy(input [][]rune) [][]rune {
	var tmp [][]rune
	for _, line := range input {
		tmpL := make([]rune, len(line))
		copy(tmpL, line)
		tmp = append(tmp, tmpL)
	}

	return tmp
}
