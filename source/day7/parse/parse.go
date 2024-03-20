package parse

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Hand string
	Bid  int
}

func Parse(path string) []Input {
	file, err := os.Open(path)
	if err != nil {
		panic(path)
	}
	defer file.Close()

	input := bufio.NewScanner(file)

	inputs := make([]Input, 0)

	for input.Scan() {
		fields := strings.Fields(input.Text())
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(input.Text())
		}
		inputs = append(inputs, Input{fields[0], bid})
	}

	if input.Err() != nil {
		panic(input.Err())
	}

	return inputs
}
