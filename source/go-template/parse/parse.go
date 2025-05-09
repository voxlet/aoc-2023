package parse

import (
	"bufio"
	"os"
	"strings"
)

type Input struct {
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
		fields := strings.Fields(lines.Text())
	}

	if lines.Err() != nil {
		panic(lines.Err())
	}

	return input
}
