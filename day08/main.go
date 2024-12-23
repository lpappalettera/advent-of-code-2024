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
	antennas := make(map[rune][]coor)
	var antinodes []coor

	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], coor{y, x})
			}
		}
	}

	for _, coors := range antennas {
		for _, a := range coors {
			for _, b := range coors {
				if a != b {
					ant1, ant2 := getAntinodes(a, b, 2)
					if canAntinote(input, ant1.y, ant1.x) {
						antinodes = appendUnique(antinodes, ant1)
					}
					if canAntinote(input, ant2.y, ant2.x) {
						antinodes = appendUnique(antinodes, ant2)
					}
				}
			}
		}
	}

	return len(antinodes)
}

func part2(filename string) int {
	input := parse(filename)
	antennas := make(map[rune][]coor)
	var antinodes []coor

	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], coor{y, x})
			}
		}
	}

	for _, coors := range antennas {
		for _, a := range coors {
			for _, b := range coors {
				if a != b {
					ant1, ant2 := getAntinodes(a, b, 2)
					if canAntinote(input, ant1.y, ant1.x) {
						antinodes = appendUnique(antinodes, ant1)
					}
					if canAntinote(input, ant2.y, ant2.x) {
						antinodes = appendUnique(antinodes, ant2)
					}
				}
			}
		}
	}

	return len(antinodes)
}

type coor struct {
	y int
	x int
}

func appendUnique(l []coor, value coor) []coor {
	contains := false
	for _, a := range l {
		if a.y == value.y && a.x == value.x {
			contains = true
		}
	}

	if !contains {
		return append(l, value)
	}

	return l
}

func getAntinodes(a coor, b coor, r int) (coor, coor) {
	var ant1 coor = coor{r*a.y - (r-1)*b.y, r*a.x - (r-1)*b.x}
	var ant2 coor = coor{r*b.y - (r-1)*a.y, r*b.x - (r-1)*a.x}
	return ant1, ant2
}

func canAntinote(input [][]rune, y int, x int) bool {
	return y >= 0 && x >= 0 && y < len(input) && x < len(input[0])
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
