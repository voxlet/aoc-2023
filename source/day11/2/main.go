package main

import (
	"github.com/voxlet/aoc-2023/source/day11/parse"
)

type Pos struct {
	row int
	col int
}

func absDist(a int, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func emptyCountIn(empties []bool, a, b int) int {
	count := 0

	start, end := minMax(a, b)
	for _, empty := range empties[start:end] {
		if empty {
			count++
		}
	}

	return count
}

func manhattanDist(a Pos, b Pos, rowEmpty, colEmpty []bool) int {
	dist := absDist(a.row, b.row) + absDist(a.col, b.col)

	expansions := 0
	expansions += emptyCountIn(rowEmpty, a.row, b.row)
	expansions += emptyCountIn(colEmpty, a.col, b.col)

	dist += expansions * (1000000 - 1)

	return dist
}

func distanceSum(galaxies []Pos, rowEmpty, colEmpty []bool) int {
	sum := 0

	for a := range len(galaxies) {
		for b := a; b < len(galaxies); b++ {
			sum += manhattanDist(galaxies[a], galaxies[b], rowEmpty, colEmpty)
		}
	}

	return sum
}

func makeAllEmpty(size int) []bool {
	empties := make([]bool, size)
	for i := range size {
		empties[i] = true
	}
	return empties
}

func main() {
	inputs := parse.Parse("source/day11/input.txt")

	rowEmpty := makeAllEmpty(len(inputs))
	colEmpty := makeAllEmpty(len(inputs[0]))
	galaxies := make([]Pos, 0)

	for row, input := range inputs {
		for col, r := range input {
			if r != '#' {
				continue
			}
			rowEmpty[row] = false
			colEmpty[col] = false

			galaxy := Pos{row, col}
			galaxies = append(galaxies, galaxy)
		}
	}

	sum := distanceSum(galaxies, rowEmpty, colEmpty)

	println(sum)
}
