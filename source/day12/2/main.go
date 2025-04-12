package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/voxlet/aoc-2023/source/day12/parse"
)

func arrangementCount(input parse.Input, known map[string]int) int {
	if input.Record == "" {
		if len(input.GroupCounts) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(input.GroupCounts) == 0 {
		if strings.IndexByte(input.Record, '#') == -1 {
			return 1
		} else {
			return 0
		}
	}

	key := fmt.Sprintf("%v", input)

	if count, ok := known[key]; ok {
		return count
	}

	count := 0
	c := input.Record[0]

	if c != '#' {
		count += arrangementCount(parse.Input{
			Record:      input.Record[1:],
			GroupCounts: input.GroupCounts,
		}, known)
	}

	if c != '.' {
		need := input.GroupCounts[0]
		recordLen := len(input.Record)

		if recordLen >= need &&
			strings.IndexByte(input.Record[1:need], '.') == -1 &&
			(recordLen == need || input.Record[need] != '#') {
			count += arrangementCount(parse.Input{
				Record:      input.Record[min(recordLen, need+1):],
				GroupCounts: input.GroupCounts[1:],
			}, known)
		}
	}

	known[key] = count

	return count
}

func repeat[T any](x T, n int) []T {
	repeated := make([]T, n)
	for i := range n {
		repeated[i] = x
	}
	return repeated
}

func unfold(inputs []parse.Input) {
	for i := range inputs {
		input := &inputs[i]
		input.Record = strings.Join(repeat(input.Record, 5), "?")
		input.GroupCounts = slices.Concat(repeat(input.GroupCounts, 5)...)
	}
}

func main() {
	inputs := parse.Parse("source/day12/input.txt")
	unfold(inputs)

	sum := 0
	for _, input := range inputs {
		sum += arrangementCount(input, make(map[string]int))
	}

	fmt.Println(sum)
}
