package parse

import (
	"bufio"
	"os"
)

type Input struct {
	Rows     []string
	ColCount int
}

func Parse(path string) Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	input := Input{
		Rows: make([]string, 0),
	}

	for lines.Scan() {
		input.Rows = append(input.Rows, lines.Text())
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	input.ColCount = len(input.Rows[0])

	return input
}
