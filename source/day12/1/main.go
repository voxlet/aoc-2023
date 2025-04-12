package main

import (
	"fmt"
	"strings"

	"github.com/voxlet/aoc-2023/source/day12/parse"
)

func valid(input parse.Input) bool {
	groups := strings.Split(input.Record, ".")

	i := 0
	for _, group := range groups {
		count := len(group)

		if count == 0 {
			continue
		}

		if i >= len(input.GroupCounts) {
			return false
		}

		if count != input.GroupCounts[i] {
			return false
		}
		i++
	}

	return i == len(input.GroupCounts)
}

func arrangementCount(input parse.Input) int {
	if strings.IndexByte(input.Record, '?') == -1 {
		if valid(input) {
			return 1
		} else {
			return 0
		}
	}

	return arrangementCount(parse.Input{
		Record:      strings.Replace(input.Record, "?", ".", 1),
		GroupCounts: input.GroupCounts,
	}) +
		arrangementCount(parse.Input{
			Record:      strings.Replace(input.Record, "?", "#", 1),
			GroupCounts: input.GroupCounts,
		})
}

func main() {
	inputs := parse.Parse("source/day12/input.txt")

	sum := 0
	for _, input := range inputs {
		sum += arrangementCount(input)
	}

	fmt.Println(inputs)
	fmt.Println(sum)
}
