package parse

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Record      string
	GroupCounts []int
}

func Parse(path string) []Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	inputs := make([]Input, 0)

	for lines.Scan() {
		fields := strings.Fields(lines.Text())

		record := fields[0]

		groupCountStrings := strings.Split(fields[1], ",")
		groupCounts := make([]int, 0, len(groupCountStrings))

		for _, s := range groupCountStrings {
			groupCount, err := strconv.Atoi(s)
			if err != nil {
				panic("bad input: " + lines.Text())
			}
			groupCounts = append(groupCounts, groupCount)
		}
		inputs = append(inputs, Input{record, groupCounts})
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return inputs
}
