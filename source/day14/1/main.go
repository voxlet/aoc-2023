package main

import (
	"fmt"

	"github.com/voxlet/aoc-2023/source/day14/parse"
)

func addLoads(rockCount int, end int, loads []int) {
	for i := range rockCount {
		row := end + 1 + i
		loads[row]++
	}
}

func main() {
	inputs := parse.Parse("source/day14/input.txt")

	loads := make([]int, len(inputs))
	colCount := len(inputs[0])

	for col := range colCount {
		rockCount := 0

		for row := len(inputs) - 1; row >= 0; row-- {
			ch := inputs[row][col]

			if ch == 'O' {
				rockCount++
				continue
			}

			if ch == '#' {
				addLoads(rockCount, row, loads)
				rockCount = 0
			}
		}

		addLoads(rockCount, -1, loads)
	}

	sum := 0

	for row, load := range loads {
		sum += load * (len(loads) - row)
	}

	fmt.Println(sum)
}
