package parse

import (
	"bufio"
	"os"
)

type Input struct {
	Rows []string
}

func Parse(path string) []Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	inputs := make([]Input, 0)
	input := Input{}

	for lines.Scan() {
		if lines.Text() == "" {
			inputs = append(inputs, input)
			input = Input{}
			continue
		}

		input.Rows = append(input.Rows, lines.Text())
	}
	inputs = append(inputs, input)

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return inputs
}
