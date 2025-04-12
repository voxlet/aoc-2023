package parse

import (
	"bufio"
	"os"
)

func Parse(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)

	inputs := make([]string, 0)

	for lines.Scan() {
		inputs = append(inputs, lines.Text())
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return inputs
}
