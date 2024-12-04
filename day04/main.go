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
	input := parse(filename)
	result := 0

	for y, row := range input {
		for x, it := range row {
			if it == 'X' {
				var m []Coor
				adjacent := []Coor{
					{y - 1, x - 1},
					{y - 1, x},
					{y - 1, x + 1},
					{y, x - 1},
					{y, x + 1},
					{y + 1, x - 1},
					{y + 1, x},
					{y + 1, x + 1},
				}

				for _, c := range adjacent {
					if inBouds(input, c.y, c.x) && input[c.y][c.x] == 'M' {
						m = append(m, c)
					}
				}

				for _, coor := range m {
					yDif := coor.y - y
					xDif := coor.x - x

					aC := Coor{y + 2*yDif, x + 2*xDif}
					sC := Coor{y + 3*yDif, x + 3*xDif}

					if inBouds(input, sC.y, sC.x) && input[aC.y][aC.x] == 'A' && input[sC.y][sC.x] == 'S' {
						result++
					}
				}

			}
		}
	}

	return result
}

func part2(filename string) int {
	input := parse(filename)
	result := 0

	for y, row := range input {
		for x, it := range row {
			if it == 'A' {

				var m []Coor
				adjacent := []Coor{
					{y - 1, x - 1},
					{y - 1, x + 1},
					{y + 1, x - 1},
					{y + 1, x + 1},
				}

				for _, c := range adjacent {
					if inBouds(input, c.y, c.x) && input[c.y][c.x] == 'M' {
						m = append(m, c)
					}
				}

				if len(m) == 2 {
					a1 := Coor{y + y - m[0].y, x + x - m[0].x}
					a2 := Coor{y + y - m[1].y, x + x - m[1].x}

					if inBouds(input, a1.y, a1.x) && inBouds(input, a2.y, a2.x) && input[a1.y][a1.x] == 'S' && input[a2.y][a2.x] == 'S' {
						result++
					}
				}

			}
		}
	}

	return result
}

func parse(filename string) [][]rune {
	lines := util.ReadLines(filename)
	var result [][]rune

	for _, line := range lines {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		result = append(result, row)
	}

	return result
}

type Coor struct {
	y int
	x int
}

func inBouds(input [][]rune, y int, x int) bool {
	return y >= 0 && x >= 0 && y < len(input) && x < len(input[0])
}
