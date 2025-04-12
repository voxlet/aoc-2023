package parse

import (
	"bufio"
	"os"
)

type Input = string

func Parse(path string) []Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	inputs := make([]Input, 0)

	for lines.Scan() {
		inputs = append(inputs, Input(lines.Text()))
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return inputs
}
