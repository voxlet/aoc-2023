package main

import (
	"github.com/voxlet/aoc-2023/source/day15/parse"
)

func hash(s string) int {
	code := 0

	for _, c := range []byte(s) {
		code += int(c)
		code *= 17
		code %= 256
	}

	return code
}

func main() {
	input := parse.Parse("source/day15/input.txt")

	sum := 0

	for _, s := range input.Steps {
		sum += hash(s)
	}

	println(sum)
}
