package parse

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	inputs := make([][]int, 0)

	for lines.Scan() {
		fields := strings.Fields(lines.Text())
		input := make([]int, 0)

		for _, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				panic("bad input: " + field + " " + err.Error())
			}
			input = append(input, n)
		}

		inputs = append(inputs, input)
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return inputs
}
