package main

import (
	"fmt"
	"slices"

	"github.com/voxlet/aoc-2023/source/day9/parse"
)

func valueFor(level int, i int, values [][]int) (int, [][]int) {
	if level < len(values) && i < len(values[level]) {
		return values[level][i], values
	}

	up := level - 1
	if up < 0 {
		fmt.Println(level, i, values[0])
		panic("out of levels")
	}

	large, values := valueFor(up, i, values)
	small, values := valueFor(up, i+1, values)

	diff := large - small

	var vs []int
	if level < len(values) {
		vs = values[level]
	} else {
		vs = make([]int, 0)
		values = append(values, vs)
	}

	values[level] = append(vs, diff)

	return diff, values
}

func allZero(vs []int) bool {
	for _, v := range vs {
		if v != 0 {
			return false
		}
	}
	return true
}

func nextValue(input []int) int {
	values := make([][]int, 0)
	values = append(values, slices.Clone(input))
	slices.Reverse(values[0])

	vs := values[0]
	level := 0

	for !allZero(vs) {
		fmt.Println(vs)
		level++

		for i := range len(vs) - 1 {
			_, values = valueFor(level, i, values)
		}
		vs = values[level]
	}

	sum := 0
	for _, vs := range values {
		sum += vs[0]
	}

	return sum
}

func main() {
	inputs := parse.Parse("source/day9/input.txt")

	fmt.Println(inputs)

	sum := 0

	for _, input := range inputs {
		value := nextValue(input)
		println("found:", value, input)

		sum += value
	}

	println("sum:", sum)
}
