package main

import (
	"slices"
	"strings"

	"github.com/voxlet/aoc-2023/source/day11/parse"
)

func emptyCol(inputs []string, col int) bool {
	for row := range len(inputs) {
		if inputs[row][col] == '#' {
			return false
		}
	}
	return true
}

func emptyRow(inputs []string, row int) bool {
	return strings.IndexByte(inputs[row], '#') == -1
}

func insertEmptyCol(inputs []string, col int) {
	for row := range len(inputs) {
		oldRow := inputs[row]
		var newRow strings.Builder

		newRow.WriteString(oldRow[:col])
		newRow.WriteByte('.')
		newRow.WriteString(oldRow[col:])

		inputs[row] = newRow.String()
	}
}

func insertEmptyRow(inputs []string, row int) []string {
	size := len(inputs[0])
	newRow := strings.Repeat(".", size)
	return slices.Insert(inputs, row, newRow)
}

func expand(inputs []string) []string {
	for col := 0; col != len(inputs[0]); col++ {
		if !emptyCol(inputs, col) {
			continue
		}
		insertEmptyCol(inputs, col)
		col++
	}
	for row := 0; row != len(inputs); row++ {
		if !emptyRow(inputs, row) {
			continue
		}
		inputs = insertEmptyRow(inputs, row)
		row++
	}
	return inputs
}

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

func manhattanDist(a Pos, b Pos) int {
	return absDist(a.row, b.row) + absDist(a.col, b.col)
}

func main() {
	inputs := parse.Parse("source/day11/input.txt")
	inputs = expand(inputs)

	galaxies := make([]Pos, 0)
	for row, input := range inputs {
		for col, r := range input {
			if r != '#' {
				continue
			}
			galaxies = append(galaxies, Pos{row, col})
		}
	}

	sum := 0
	for a := range len(galaxies) {
		for b := a; b < len(galaxies); b++ {
			sum += manhattanDist(galaxies[a], galaxies[b])
		}
	}

	println(sum)
}
