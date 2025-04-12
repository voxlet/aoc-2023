package parse

import (
	"bufio"
	"os"
)

type Input struct {
	Rows    []string
	ColSize int
}

func Parse(path string) Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	input := Input{}

	for lines.Scan() {
		input.Rows = append(input.Rows, lines.Text())
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	input.ColSize = len(input.Rows[0])

	return input
}
