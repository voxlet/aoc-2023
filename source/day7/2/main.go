package main

import (
	"cmp"
	"fmt"
	"slices"
	"unicode"

	"github.com/voxlet/aoc-2023/source/day7/parse"
)

const JOKER = 1

var valuesByLabel = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': JOKER,
	'T': 10,
}

func toValue(c rune) int {
	if unicode.IsDigit(c) {
		return int(c - '0')
	}
	v, ok := valuesByLabel[c]
	if !ok {
		panic(c)
	}
	return v
}

func score(hand string) int {
	values := make([]int, 0, len(hand))
	var counts [14]int
	var jokerCount int

	for _, c := range hand {
		v := toValue(c)
		values = append(values, v)

		if v == JOKER {
			jokerCount++
		} else {
			counts[v-1]++
		}
	}

	slices.Reverse(values)

	score := 0

	factor := 1

	for _, v := range values {
		score += v * factor
		factor *= 100
	}

	slices.Sort(counts[:])
	slices.Reverse(counts[:])

	typeScore := counts[0] + jokerCount

	if typeScore >= 4 {
		typeScore++
	}
	if typeScore >= 3 {
		typeScore++
	}
	if typeScore >= 2 && counts[1] == 2 {
		typeScore++
	}

	score += typeScore * factor

	return score
}

func main() {
	inputs := parse.Parse("source/day7/input.txt")

	slices.SortFunc(inputs, func(a parse.Input, b parse.Input) int {
		return cmp.Compare(score(a.Hand), score(b.Hand))
	})

	winnings := 0

	for i, input := range inputs {
		fmt.Println(i+1, input, score(input.Hand))
		winnings += (i + 1) * input.Bid
	}

	fmt.Println(inputs, len(inputs), winnings)
}
