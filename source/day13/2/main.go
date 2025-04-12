package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/voxlet/aoc-2023/source/day13/parse"
)

func rotate(vs []string) []string {
	count := len(vs[0])
	res := make([]string, 0, count)

	for i := range count {
		var b strings.Builder
		for n := len(vs) - 1; n >= 0; n-- {
			b.WriteByte(vs[n][i])
		}
		res = append(res, b.String())
	}

	return res
}

func diffs(a string, b string) int {
	count := 0
	for i := range len(a) {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

func find(vs []string) (int, error) {
	for i := 1; i < len(vs); i++ {
		diffSum := 0
		f, b := i, i-1

		for {
			diffSum += diffs(vs[b], vs[f])
			if diffSum > 1 {
				break
			}

			if b == 0 || f == len(vs)-1 {
				if diffSum == 1 {
					return i, nil
				}
				break
			}
			b--
			f++
		}
	}

	return -1, fmt.Errorf("not found")
}

func summary(input parse.Input) int {
	if found, err := find(input.Rows); err == nil {
		return found * 100
	}
	cols := rotate(input.Rows)
	if found, err := find(cols); err == nil {
		return found
	}
	j, _ := json.MarshalIndent(input, "", "  ")
	panic(fmt.Sprintf("not found: %s", string(j)))
}

func main() {
	inputs := parse.Parse("source/day13/input.txt")
	sum := 0

	for _, input := range inputs {
		s := summary(input)
		j, _ := json.MarshalIndent(input, "", "  ")
		fmt.Printf("%v: ", s)
		println(string(j))
		sum += s
	}

	println(sum)
}
