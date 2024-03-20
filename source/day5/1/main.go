package main

import (
	"fmt"
	"slices"

	"github.com/voxlet/aoc-2023/source/day5/mapping"
)

func mapSeeds(seeds []int, mappings [][]mapping.Entry) []int {
	locations := make([]int, 0)

	for _, v := range seeds {
		for _, m := range mappings {
			v = mapping.Apply(v, m)
		}
		locations = append(locations, v)
	}

	return locations
}

func main() {
	seeds, mappings := mapping.Parse("source/day5/input.txt")
	locations := mapSeeds(seeds, mappings)

	fmt.Println("min:", slices.Min(locations))
}
